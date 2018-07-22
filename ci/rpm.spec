Name:           %{name}
Url:            https://github.com/openstack/%{name}
Summary:        %{name}
License:        Apache-2.0
Group:          System/Emulators/PC
Version:        %{version}
Release:        %{release}
Source0:        src.tar.gz
BuildRoot:     %{_tmppath}/%{name}-%{version}-%{release}
AutoReq: no

%description
%{name}

%prep
rm -rf %{buildroot}
tar -xf ../SOURCES/src.tar.gz

%build
ls -al
mkdir -p opt/%{name}/bin
cp src/go-sample-webapp opt/%{name}/bin


%install
rm -rf %{buildroot}
mkdir -p %{buildroot}
mkdir -p %{buildroot}/etc
mkdir -p %{buildroot}/var/log/%{name}
mkdir -p %{buildroot}/usr/lib/systemd
cp -r opt %{buildroot}
cp -r src/base/ci/system %{buildroot}/usr/lib/systemd/system
cp -r src/base/ci/etc %{buildroot}/etc/%{name}


%clean
rm -rf %{buildroot}


%files
/opt/%{name}
%attr(-, root, root) /usr/lib/systemd/system/*
%dir %attr(0755, root, root) /var/log/%{name}
%dir %attr(0755, root, root) /etc/%{name}
%config(noreplace) %attr(-, root, root) /etc/%{name}/*


%changelog
