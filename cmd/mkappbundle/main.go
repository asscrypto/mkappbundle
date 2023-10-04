package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	//	"html/template"
	"io"
	"log"
	"os"
	"path"
	"text/template"
)

const singleBinaryInfoTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>CFBundleExecutable</key>
  <string>{{ .AppName }}</string>
  <key>CFBundleIconFile</key>
  <string>{{ .AppName }}.icns</string>
  <key>CFBundleInfoDictionaryVersion</key>
  <string>6.0</string>
  <key>CFBundlePackageType</key>
  <string>APPL</string>
  <key>CFBundleVersion</key>
  <string>{{ .AppVersion }}</string>
  <key>NSHighResolutionCapable</key>
  <string>True</string>
</dict>
</plist>
`

var (
	buildDir    = "buildd"
	pkgBuildDir = "pkgbuilddir"

	appName    string
	appVersion string
)

func main() {
	if envvar := os.Getenv("PKGBUILDDIR"); envvar != "" {
		pkgBuildDir = envvar
	}

	if envvar := os.Getenv("BUILDDIR"); envvar != "" {
		buildDir = envvar
	}

	appName, appVersion = os.Getenv("APP_NAME"), os.Getenv("APP_VERSION")
	if appName == "" || appVersion == "" {
		log.Fatal("the environment variables APP_NAME and APP_VERSION cannot be nil.")
	}

	appNameTitle := cases.Title(language.English).String(appName)
	appBundleDir := path.Join(pkgBuildDir, fmt.Sprintf("%s.app", appNameTitle))
	macOSDir := path.Join(appBundleDir, "Contents", "MacOS")

	if err := os.MkdirAll(macOSDir, 0750); err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(path.Join(appBundleDir, "Contents", "Resources"), 0750); err != nil {
		log.Fatal(err)
	}

	src, dst := path.Join(buildDir, appName), path.Join(macOSDir, appName)
	log.Printf("copying %q -> %q:", src, dst)
	if err := copyFile(src, dst); err != nil {
		log.Fatalf("couldn't copy the file: %v", err)
	}

	//t := template.Must(template.ParseFiles([]string{"/Users/alessio/devel/mkappbundle/app_template.txt"}...))
	t := template.Must(template.New("plist").Parse(singleBinaryInfoTemplate))

	data := struct {
		AppName    string
		AppVersion string
	}{
		AppName:    appName,
		AppVersion: appVersion,
	}

	//w := new(bytes.Buffer)
	//if err := xml.EscapeText(f, w.Bytes()); err != nil {
	//	log.Fatal(err)
	//}
	//
	//xmltemplate.HTMLEscape()

	f, err := os.Create(path.Join(appBundleDir, "Contents", "Info.plist"))
	if err != nil {
		log.Fatal(err)
	}

	defer fileCloser(f)

	if err := t.Execute(f, data); err != nil {
		log.Fatal(err)
	}

}

//func checkError(err error, v interface{}) {
//	if err == nil {
//		return
//	}
//
//	if v != nil {
//		log.Fatalf("%v: %v", v, err)
//	}
//
//	log.Fatal(err)
//
//	//switch len(ss) {
//	//case 0:
//	//	log.Fatal(err)
//	//case 1:
//	//	log.Fatalf("%s: %v")
//	//default:
//	//	log.Fatalf("%s: %v", strings.Join(ss, ":"), err)
//	//}
//	//
//	//if len(ss) == 0 {
//	//	log.Fatal(err)
//	//}
//
//}

func fileCloser(f io.Closer) {
	if err := f.Close(); err != nil {
		log.Fatalf("error while closing the file: %v", err)
	}
}

func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fileCloser(source)

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer fileCloser(destination)

	_, err = io.Copy(destination, source)

	return err
}
