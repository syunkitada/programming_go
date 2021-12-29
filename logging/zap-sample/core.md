# コア部分

```go:logger.go
 57 // New constructs a new Logger from the provided zapcore.Core and Options. If
 58 // the passed zapcore.Core is nil, it falls back to using a no-op
 59 // implementation.
 60 //
 61 // This is the most flexible way to construct a Logger, but also the most
 62 // verbose. For typical use cases, the highly-opinionated presets
 63 // (NewProduction, NewDevelopment, and NewExample) or the Config struct are
 64 // more convenient.
 65 //
 66 // For sample code, see the package-level AdvancedConfiguration example.
 67 func New(core zapcore.Core, options ...Option) *Logger {
 68     if core == nil {
 69         return NewNop()
 70     }
 71     log := &Logger{
 72         core:        core,
 73         errorOutput: zapcore.Lock(os.Stderr),
 74         addStack:    zapcore.FatalLevel + 1,
 75         clock:       zapcore.DefaultClock,
 76     }
 77     return log.WithOptions(options...)
 78 }
```

```go:logger.go
188 // Info logs a message at InfoLevel. The message includes any fields passed
189 // at the log site, as well as any fields accumulated on the logger.
190 func (log *Logger) Info(msg string, fields ...Field) {
191     if ce := log.check(InfoLevel, msg); ce != nil {
192         ce.Write(fields...)
193     }
194 }
```

```go:logger.go
261 func (log *Logger) check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
262     // check must always be called directly by a method in the Logger interface
263     // (e.g., Check, Info, Fatal).
264     const callerSkipOffset = 2
265
266     // Check the level first to reduce the cost of disabled log calls.
267     // Since Panic and higher may exit, we skip the optimization for those levels.
268     if lvl < zapcore.DPanicLevel && !log.core.Enabled(lvl) {
269         return nil
270     }
271
272     // Create basic checked entry thru the core; this will be non-nil if the
273     // log message will actually be written somewhere.
274     ent := zapcore.Entry{
275         LoggerName: log.name,
276         Time:       log.clock.Now(),
277         Level:      lvl,
278         Message:    msg,
279     }
280     ce := log.core.Check(ent, nil)
281     willWrite := ce != nil
282
283     // Set up any required terminal behavior.
284     switch ent.Level {
285     case zapcore.PanicLevel:
286         ce = ce.Should(ent, zapcore.WriteThenPanic)
287     case zapcore.FatalLevel:
288         onFatal := log.onFatal
289         // Noop is the default value for CheckWriteAction, and it leads to
290         // continued execution after a Fatal which is unexpected.
291         if onFatal == zapcore.WriteThenNoop {
292             onFatal = zapcore.WriteThenFatal
293         }
294         ce = ce.Should(ent, onFatal)
295     case zapcore.DPanicLevel:
296         if log.development {
297             ce = ce.Should(ent, zapcore.WriteThenPanic)
298         }
299     }
300
301     // Only do further annotation if we're going to write this message; checked
302     // entries that exist only for terminal behavior don't benefit from
303     // annotation.
304     if !willWrite {
305         return ce
306     }
307
308     // Thread the error output through to the CheckedEntry.
309     ce.ErrorOutput = log.errorOutput
310     if log.addCaller {
311         frame, defined := getCallerFrame(log.callerSkip + callerSkipOffset)
312         if !defined {
313             fmt.Fprintf(log.errorOutput, "%v Logger.check error: failed to get caller\n", ent.Time.UTC())
314             log.errorOutput.Sync()
315         }
316
317         ce.Entry.Caller = zapcore.EntryCaller{
318             Defined:  defined,
319             PC:       frame.PC,
320             File:     frame.File,
321             Line:     frame.Line,
322             Function: frame.Function,
323         }
324     }
325     if log.addStack.Enabled(ce.Entry.Level) {
326         ce.Entry.Stack = StackSkip("", log.callerSkip+callerSkipOffset).String
327     }
328
329     return ce
330 }
```

```go:zapcore/entry.go
# logのentryは使いまわされる
# checkedは使い終わったやつ
 36 var (
 37     _cePool = sync.Pool{New: func() interface{} {
 38         // Pre-allocate some space for cores.
 39         return &CheckedEntry{
 40             cores: make([]Core, 4),
 41         }
 42     }}
 43 )
 44
 45 func getCheckedEntry() *CheckedEntry {
 46     ce := _cePool.Get().(*CheckedEntry)
 47     ce.reset()
 48     return ce
 49 }
 50
 51 func putCheckedEntry(ce *CheckedEntry) {
 52     if ce == nil {
 53         return
 54     }
 55     _cePool.Put(ce)
 56 }

171 // CheckedEntry is an Entry together with a collection of Cores that have
172 // already agreed to log it.
173 //
174 // CheckedEntry references should be created by calling AddCore or Should on a
175 // nil *CheckedEntry. References are returned to a pool after Write, and MUST
176 // NOT be retained after calling their Write method.
177 type CheckedEntry struct {
178     Entry
179     ErrorOutput WriteSyncer
180     dirty       bool // best-effort detection of pool misuse
181     should      CheckWriteAction
182     cores       []Core
183 }


197 // Write writes the entry to the stored Cores, returns any errors, and returns
198 // the CheckedEntry reference to a pool for immediate re-use. Finally, it
199 // executes any required CheckWriteAction.
200 func (ce *CheckedEntry) Write(fields ...Field) {
201     if ce == nil {
202         return
203     }
204
205     if ce.dirty {
206         if ce.ErrorOutput != nil {
207             // Make a best effort to detect unsafe re-use of this CheckedEntry.
208             // If the entry is dirty, log an internal error; because the
209             // CheckedEntry is being used after it was returned to the pool,
210             // the message may be an amalgamation from multiple call sites.
211             fmt.Fprintf(ce.ErrorOutput, "%v Unsafe CheckedEntry re-use near Entry %+v.\n", ce.Time, ce.Entry)
212             ce.ErrorOutput.Sync()
213         }
214         return
215     }
216     ce.dirty = true
217
218     var err error
219     for i := range ce.cores {
// ここでログが書き込まれる
220         err = multierr.Append(err, ce.cores[i].Write(ce.Entry, fields))
221     }
222     if err != nil && ce.ErrorOutput != nil {
223         fmt.Fprintf(ce.ErrorOutput, "%v write error: %v\n", ce.Time, err)
224         ce.ErrorOutput.Sync()
225     }
226
227     should, msg := ce.should, ce.Message
228     putCheckedEntry(ce)
229
230     switch should {
231     case WriteThenPanic:
232         panic(msg)
233     case WriteThenFatal:
234         exit.Exit()
235     case WriteThenGoexit:
236         runtime.Goexit()
237     }
238 }
```

```go:zapcore/core.go
 85 func (c *ioCore) Write(ent Entry, fields []Field) error {
 86     buf, err := c.enc.EncodeEntry(ent, fields)
 87     if err != nil {
 88         return err
 89     }
 90     _, err = c.out.Write(buf.Bytes())
 91     buf.Free()
 92     if err != nil {
 93         return err
 94     }
 95     if ent.Level > ErrorLevel {
 96         // Since we may be crashing the program, sync the output. Ignore Sync
 97         // errors, pending a clean solution to issue #370.
 98         c.Sync()
 99     }
100     return nil
101 }
```
