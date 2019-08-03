package server

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/syunkitada/go-samples/generator-sample/pkg/v1"
	v1resolver "github.com/syunkitada/go-samples/generator-sample/pkg/v1/resolver"

	v2 "github.com/syunkitada/go-samples/generator-sample/pkg/v2"
	v2resolver "github.com/syunkitada/go-samples/generator-sample/pkg/v2/resolver"
)

func Serv() {
	r := gin.Default()

	v1Resolver := v1resolver.New()
	v1.New(r, v1Resolver)

	v2Resolver := v2resolver.New()
	v2.New(r, v2Resolver)

	r.Run()
}
