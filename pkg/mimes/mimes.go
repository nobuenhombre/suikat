package mimes

// https://developer.mozilla.org/ru/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types

const (
	AACAudio                          = "audio/aac"
	AbiWordDocument                   = "application/x-abiword"
	ArchiveDocument                   = "application/x-freearc"
	AudioVideoInterleave              = "video/x-msvideo"
	AmazonKindleEBook                 = "application/vnd.amazon.ebook"
	BinaryData                        = "application/octet-stream"
	WindowsBitmapGraphics             = "image/bmp"
	BZipArchive                       = "application/x-bzip"
	BZip2Archive                      = "application/x-bzip2"
	CShellScript                      = "application/x-csh"
	CascadingStyleSheets              = "text/css"
	CommaSeparatedValues              = "text/csv"
	MicrosoftWord                     = "application/msword"
	MicrosoftWordOpenXML              = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	MicrosoftExcel                    = "application/vnd.ms-excel"
	MicrosoftExcelOpenXML             = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	MicrosoftPowerPoint               = "application/vnd.ms-powerpoint"
	MicrosoftPowerPointOpenXML        = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	MicrosoftVisio                    = "application/vnd.visio"
	MicrosoftEmbeddedOpenTypeFont     = "application/vnd.ms-fontobject"
	ElectronicPublication             = "application/epub+zip"
	GZipCompressedArchive             = "application/gzip"
	GraphicsInterchangeFormat         = "image/gif"
	HyperTextMarkupLanguage           = "text/html"
	Icon                              = "image/vnd.microsoft.icon"
	ICalendar                         = "text/calendar"
	JavaArchive                       = "application/java-archive"
	JPEGImages                        = "image/jpeg"
	JavaScript                        = "text/javascript"
	JSON                              = "application/json"
	JSONLD                            = "application/ld+json"
	MusicalInstrumentDigitalInterface = "audio/midi"
	MP3Audio                          = "audio/mpeg"
	MPEGVideo                         = "video/mpeg"
	AppleInstallerPackage             = "application/vnd.apple.installer+xml"
	OpenDocumentPresentation          = "application/vnd.oasis.opendocument.presentation"
	OpenDocumentSpreadsheet           = "application/vnd.oasis.opendocument.spreadsheet"
	OpenDocumentText                  = "application/vnd.oasis.opendocument.text"
	OGGAudio                          = "audio/ogg"
	OGGVideo                          = "video/ogg"
	OGG                               = "application/ogg"
	OpusAudio                         = "audio/opus"
	OpenTypeFont                      = "font/otf"
	PortableNetworkGraphics           = "image/png"
	AdobePortableDocumentFormat       = "application/pdf"
	PHPHypertextPreprocessor          = "application/php"
	RARArchive                        = "application/vnd.rar"
	RichTextFormat                    = "application/rtf"
	BourneShellScript                 = "application/x-sh"
	ScalableVectorGraphics            = "image/svg+xml"
	AdobeFlashDocument                = "application/x-shockwave-flash"
	TapeArchive                       = "application/x-tar"
	TaggedImageFileFormat             = "image/tiff"
	MPEGTransportStream               = "video/mp2t"
	TrueTypeFont                      = "font/ttf"
	Text                              = "text/plain"
	WaveformAudio                     = "audio/wav"
	WEBMAudio                         = "audio/webm"
	WEBMVideo                         = "video/webm"
	WEBPImage                         = "image/webp"
	WebOpenFontFormat                 = "font/woff"
	WebOpenFontFormat2                = "font/woff2"
	XHTML                             = "application/xhtml+xml"
	ZIPArchive                        = "application/zip"
	SevenZipArchive                   = "application/x-7z-compressed"
	XML                               = "application/xml"
)

func mimesByExt() map[string]string {
	return map[string]string{
		".aac":    AACAudio,
		".abw":    AbiWordDocument,
		".arc":    ArchiveDocument,
		".avi":    AudioVideoInterleave,
		".azw":    AmazonKindleEBook,
		".bin":    BinaryData,
		".bmp":    WindowsBitmapGraphics,
		".bz":     BZipArchive,
		".bz2":    BZip2Archive,
		".csh":    CShellScript,
		".css":    CascadingStyleSheets,
		".csv":    CommaSeparatedValues,
		".doc":    MicrosoftWord,
		".docx":   MicrosoftWordOpenXML,
		".eot":    MicrosoftEmbeddedOpenTypeFont,
		".epub":   ElectronicPublication,
		".gz":     GZipCompressedArchive,
		".gif":    GraphicsInterchangeFormat,
		".htm":    HyperTextMarkupLanguage,
		".html":   HyperTextMarkupLanguage,
		".ico":    Icon,
		".ics":    ICalendar,
		".jar":    JavaArchive,
		".jpeg":   JPEGImages,
		".jpg":    JPEGImages,
		".js":     JavaScript,
		".mjs":    JavaScript,
		".json":   JSON,
		".jsonld": JSONLD,
		".mid":    MusicalInstrumentDigitalInterface,
		".midi":   MusicalInstrumentDigitalInterface,
		".mp3":    MP3Audio,
		".mpeg":   MPEGVideo,
		".mpkg":   AppleInstallerPackage,
		".odp":    OpenDocumentPresentation,
		".ods":    OpenDocumentSpreadsheet,
		".odt":    OpenDocumentText,
		".oga":    OGGAudio,
		".ogv":    OGGVideo,
		".ogx":    OGG,
		".opus":   OpusAudio,
		".otf":    OpenTypeFont,
		".png":    PortableNetworkGraphics,
		".pdf":    AdobePortableDocumentFormat,
		".php":    PHPHypertextPreprocessor,
		".ppt":    MicrosoftPowerPoint,
		".pptx":   MicrosoftPowerPointOpenXML,
		".rar":    RARArchive,
		".rtf":    RichTextFormat,
		".sh":     BourneShellScript,
		".svg":    ScalableVectorGraphics,
		".swf":    AdobeFlashDocument,
		".tar":    TapeArchive,
		".tif":    TaggedImageFileFormat,
		".tiff":   TaggedImageFileFormat,
		".ts":     MPEGTransportStream,
		".ttf":    TrueTypeFont,
		".txt":    Text,
		".vsd":    MicrosoftVisio,
		".wav":    WaveformAudio,
		".weba":   WEBMAudio,
		".webm":   WEBMVideo,
		".webp":   WEBPImage,
		".woff":   WebOpenFontFormat,
		".woff2":  WebOpenFontFormat2,
		".xhtml":  XHTML,
		".xls":    MicrosoftExcel,
		".xlsx":   MicrosoftExcelOpenXML,
		".zip":    ZIPArchive,
		".7z":     SevenZipArchive,
		".xml":    XML,
	}
}

func GetByExt(ext string) string {
	list := mimesByExt()

	mime, found := list[ext]
	if found {
		return mime
	}

	return BinaryData
}
