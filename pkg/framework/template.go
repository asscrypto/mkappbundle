package framework

import "text/template"

var t *template.Template

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

func init() {
	t = template.Must(template.New("frameworkPlist").
		Parse(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>CFBundleName</key>
  <string>{{ .AppName }}</string>
  <key>CFBundleDisplayName</key>
  <string>{{ .AppName }}</string>
  <key>CFBundleIdentifier</key>
  <string>{{ .Identifier }}.icns</string>
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
`))

}
