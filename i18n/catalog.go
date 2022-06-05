// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
		"es_ES": &dictionary{index: es_ESIndex, data: es_ESData},
		"ja_JP": &dictionary{index: ja_JPIndex, data: ja_JPData},
		"ko_KR": &dictionary{index: ko_KRIndex, data: ko_KRData},
		"zh_CN": &dictionary{index: zh_CNIndex, data: zh_CNData},
		"zh_TW": &dictionary{index: zh_TWIndex, data: zh_TWData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"* Leave blank to record no log and delete the original log file.": 54,
	"* Refer to the [common] section of the FRP configuration file.":   93,
	"* Refer to the parameters supported by FRP.":                      136,
	"About":         10,
	"Add":           152,
	"Admin":         59,
	"Admin Address": 60,
	"Admin Port":    61,
	"Advanced":      84,
	"All Files":     4,
	"An error occurred while checking for a software update.": 16,
	"Another config already exists with the name \"%s\".":     33,
	"Are you sure you would like to delete config \"%s\"?":    36,
	"Are you sure you would like to delete proxy \"%s\"?":     162,
	"Are you sure you would like to disable proxy \"%s\"?":    164,
	"Assets":                     63,
	"Audience":                   48,
	"Auth":                       43,
	"Auth Method":                44,
	"Authentication":             50,
	"Bandwidth":                  113,
	"Basic":                      38,
	"Bind Address":               105,
	"Bind Port":                  106,
	"Built on: %s":               3,
	"Cancel":                     19,
	"Certificate":                78,
	"Certificate Files":          6,
	"Certificate Key":            80,
	"Check Interval":             135,
	"Check Timeout":              133,
	"Check Type":                 132,
	"Check for updates":          13,
	"Checking for updates":       12,
	"Compression":                117,
	"Config already exists":      94,
	"Configuration":              20,
	"Configuration Files":        5,
	"Connection":                 66,
	"Copy":                       147,
	"Copy Access Address":        160,
	"Custom":                     91,
	"Custom Domains":             109,
	"Custom Options":             92,
	"Debug":                      65,
	"Delete":                     29,
	"Delete config \"%s\"":       35,
	"Delete proxy \"%s\"":        161,
	"Dial Timeout":               70,
	"Direct Edit":                159,
	"Disable":                    165,
	"Disable auto-start at boot": 90,
	"Disable proxy \"%s\"":       163,
	"Download updates":           11,
	"Edit":                       23,
	"Edit Config":                37,
	"Edit Proxy":                 96,
	"Empty":                      115,
	"Enable":                     157,
	"Encryption":                 116,
	"Error":                      168,
	"Exit after login failure":   89,
	"Export All Configs to ZIP":  28,
	"FRP Manager":                0,
	"FRP version: %s":            2,
	"Failure Count":              134,
	"For FRP configuration documentation, please visit the FRP project page:": 15,
	"For comments or to report bugs, please visit the project page:":          14,
	"Group":                                  129,
	"Group Key":                              130,
	"HTTP File Server":                       155,
	"HTTP Password":                          119,
	"HTTP Proxy":                             68,
	"HTTP User":                              118,
	"Health Check":                           131,
	"Heart Beats":                            51,
	"Heartbeat Interval":                     73,
	"Heartbeat Timeout":                      74,
	"Host Name":                              77,
	"Host Rewrite":                           120,
	"Import Config":                          32,
	"Import Config from File":                22,
	"Import from Clipboard":                  27,
	"Import from File":                       26,
	"Key Files":                              7,
	"Level":                                  57,
	"Load Balance":                           128,
	"Local Address":                          102,
	"Local Directory":                        151,
	"Local Path":                             125,
	"Local Port":                             103,
	"Locations":                              110,
	"Log":                                    53,
	"Log File":                               55,
	"Log Files":                              8,
	"Manual Settings":                        31,
	"Max Days":                               58,
	"Multiplexer":                            111,
	"Mux Keepalive":                          86,
	"Name":                                   39,
	"New Config":                             30,
	"New Configuration":                      21,
	"New Version!":                           9,
	"None":                                   45,
	"OK":                                     18,
	"Off":                                    76,
	"On":                                     75,
	"Open Config":                            158,
	"Open File":                              24,
	"Open Log Folder":                        139,
	"Other Options":                          88,
	"Passive Port Range":                     167,
	"Password":                               62,
	"Plugin":                                 121,
	"Plugin Name":                            122,
	"Pool Count":                             69,
	"Port":                                   166,
	"Protocol":                               67,
	"Proxy Version":                          114,
	"Proxy already exists":                   137,
	"Quick Add":                              153,
	"Random":                                 97,
	"Remote Address":                         146,
	"Remote Desktop":                         154,
	"Remote Port":                            104,
	"Role":                                   99,
	"Route User":                             112,
	"Running":                                141,
	"SOCKS5 Proxy":                           156,
	"Secret":                                 47,
	"Secret Key":                             101,
	"Select Certificate File":                79,
	"Select Certificate Key File":            81,
	"Select Log File":                        56,
	"Select Trusted CA File":                 83,
	"Select Unix Path":                       124,
	"Select a folder for directory listing.": 126,
	"Select a local directory that the admin server will load resources from.": 64,
	"Server Address":                         40,
	"Server Name":                            107,
	"Server Port":                            41,
	"Service":                                149,
	"Show in Folder":                         25,
	"Source Address":                         87,
	"Start":                                  148,
	"Starting":                               143,
	"Status":                                 145,
	"Stop":                                   150,
	"Stopped":                                142,
	"Stopping":                               144,
	"Strip Prefix":                           127,
	"Subdomain":                              108,
	"TCP Keep-Alive":                         72,
	"TCP Mux":                                85,
	"The config name \"%s\" already exists.": 95,
	"The proxy name \"%s\" already exists.":  138,
	"There are currently no updates available.": 17,
	"Token":                       46,
	"Token Endpoint":              49,
	"Trusted CA":                  82,
	"Type":                        98,
	"Unable to copy file \"%s\".": 34,
	"Unix Path":                   123,
	"Unknown":                     140,
	"User":                        42,
	"Version: %s":                 1,
	"Visitor":                     100,
	"Work Conns":                  52,
	"s":                           71,
}

var en_USIndex = []uint32{ // 170 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000c, 0x0000001b, 0x0000002e,
	0x0000003e, 0x00000048, 0x0000005c, 0x0000006e,
	0x00000078, 0x00000082, 0x0000008f, 0x00000095,
	0x000000a6, 0x000000bb, 0x000000cd, 0x0000010c,
	0x00000154, 0x0000018c, 0x000001b6, 0x000001b9,
	0x000001c0, 0x000001ce, 0x000001e0, 0x000001f8,
	0x000001fd, 0x00000207, 0x00000216, 0x00000227,
	0x0000023d, 0x00000257, 0x0000025e, 0x00000269,
	// Entry 20 - 3F
	0x00000279, 0x00000287, 0x000002bc, 0x000002d9,
	0x000002ef, 0x00000325, 0x00000331, 0x00000337,
	0x0000033c, 0x0000034b, 0x00000357, 0x0000035c,
	0x00000361, 0x0000036d, 0x00000372, 0x00000378,
	0x0000037f, 0x00000388, 0x00000397, 0x000003a6,
	0x000003b2, 0x000003bd, 0x000003c1, 0x00000402,
	0x0000040b, 0x0000041b, 0x00000421, 0x0000042a,
	0x00000430, 0x0000043e, 0x00000449, 0x00000452,
	// Entry 40 - 5F
	0x00000459, 0x000004a2, 0x000004a8, 0x000004b3,
	0x000004bc, 0x000004c7, 0x000004d2, 0x000004df,
	0x000004e1, 0x000004f0, 0x00000503, 0x00000515,
	0x00000518, 0x0000051c, 0x00000526, 0x00000532,
	0x0000054a, 0x0000055a, 0x00000576, 0x00000581,
	0x00000598, 0x000005a1, 0x000005a9, 0x000005b7,
	0x000005c6, 0x000005d4, 0x000005ed, 0x00000608,
	0x0000060f, 0x0000061e, 0x0000065d, 0x00000673,
	// Entry 60 - 7F
	0x0000069b, 0x000006a6, 0x000006ad, 0x000006b2,
	0x000006b7, 0x000006bf, 0x000006ca, 0x000006d8,
	0x000006e3, 0x000006ef, 0x000006fc, 0x00000706,
	0x00000712, 0x0000071c, 0x0000072b, 0x00000735,
	0x00000741, 0x0000074c, 0x00000756, 0x00000764,
	0x0000076a, 0x00000775, 0x00000781, 0x0000078b,
	0x00000799, 0x000007a6, 0x000007ad, 0x000007b9,
	0x000007c3, 0x000007d4, 0x000007df, 0x00000806,
	// Entry 80 - 9F
	0x00000813, 0x00000820, 0x00000826, 0x00000830,
	0x0000083d, 0x00000848, 0x00000856, 0x00000864,
	0x00000873, 0x0000089f, 0x000008b4, 0x000008db,
	0x000008eb, 0x000008f3, 0x000008fb, 0x00000903,
	0x0000090c, 0x00000915, 0x0000091c, 0x0000092b,
	0x00000930, 0x00000936, 0x0000093e, 0x00000943,
	0x00000953, 0x00000957, 0x00000961, 0x00000970,
	0x00000981, 0x0000098e, 0x00000995, 0x000009a1,
	// Entry A0 - BF
	0x000009ad, 0x000009c1, 0x000009d6, 0x00000a0b,
	0x00000a21, 0x00000a57, 0x00000a5f, 0x00000a64,
	0x00000a77, 0x00000a7d,
} // Size: 704 bytes

const en_USData string = "" + // Size: 2685 bytes
	"\x02FRP Manager\x02Version: %[1]s\x02FRP version: %[1]s\x02Built on: %[1" +
	"]s\x02All Files\x02Configuration Files\x02Certificate Files\x02Key Files" +
	"\x02Log Files\x02New Version!\x02About\x02Download updates\x02Checking f" +
	"or updates\x02Check for updates\x02For comments or to report bugs, pleas" +
	"e visit the project page:\x02For FRP configuration documentation, please" +
	" visit the FRP project page:\x02An error occurred while checking for a s" +
	"oftware update.\x02There are currently no updates available.\x02OK\x02Ca" +
	"ncel\x02Configuration\x02New Configuration\x02Import Config from File" +
	"\x02Edit\x02Open File\x02Show in Folder\x02Import from File\x02Import fr" +
	"om Clipboard\x02Export All Configs to ZIP\x02Delete\x02New Config\x02Man" +
	"ual Settings\x02Import Config\x02Another config already exists with the " +
	"name \x22%[1]s\x22.\x02Unable to copy file \x22%[1]s\x22.\x02Delete conf" +
	"ig \x22%[1]s\x22\x02Are you sure you would like to delete config \x22%[1" +
	"]s\x22?\x02Edit Config\x02Basic\x02Name\x02Server Address\x02Server Port" +
	"\x02User\x02Auth\x02Auth Method\x02None\x02Token\x02Secret\x02Audience" +
	"\x02Token Endpoint\x02Authentication\x02Heart Beats\x02Work Conns\x02Log" +
	"\x02* Leave blank to record no log and delete the original log file.\x02" +
	"Log File\x02Select Log File\x02Level\x02Max Days\x02Admin\x02Admin Addre" +
	"ss\x02Admin Port\x02Password\x02Assets\x02Select a local directory that " +
	"the admin server will load resources from.\x02Debug\x02Connection\x02Pro" +
	"tocol\x02HTTP Proxy\x02Pool Count\x02Dial Timeout\x02s\x02TCP Keep-Alive" +
	"\x02Heartbeat Interval\x02Heartbeat Timeout\x02On\x02Off\x02Host Name" +
	"\x02Certificate\x02Select Certificate File\x02Certificate Key\x02Select " +
	"Certificate Key File\x02Trusted CA\x02Select Trusted CA File\x02Advanced" +
	"\x02TCP Mux\x02Mux Keepalive\x02Source Address\x02Other Options\x02Exit " +
	"after login failure\x02Disable auto-start at boot\x02Custom\x02Custom Op" +
	"tions\x02* Refer to the [common] section of the FRP configuration file." +
	"\x02Config already exists\x02The config name \x22%[1]s\x22 already exist" +
	"s.\x02Edit Proxy\x02Random\x02Type\x02Role\x02Visitor\x02Secret Key\x02L" +
	"ocal Address\x02Local Port\x02Remote Port\x02Bind Address\x02Bind Port" +
	"\x02Server Name\x02Subdomain\x02Custom Domains\x02Locations\x02Multiplex" +
	"er\x02Route User\x02Bandwidth\x02Proxy Version\x02Empty\x02Encryption" +
	"\x02Compression\x02HTTP User\x02HTTP Password\x02Host Rewrite\x02Plugin" +
	"\x02Plugin Name\x02Unix Path\x02Select Unix Path\x02Local Path\x02Select" +
	" a folder for directory listing.\x02Strip Prefix\x02Load Balance\x02Grou" +
	"p\x02Group Key\x02Health Check\x02Check Type\x02Check Timeout\x02Failure" +
	" Count\x02Check Interval\x02* Refer to the parameters supported by FRP." +
	"\x02Proxy already exists\x02The proxy name \x22%[1]s\x22 already exists." +
	"\x02Open Log Folder\x02Unknown\x02Running\x02Stopped\x02Starting\x02Stop" +
	"ping\x02Status\x02Remote Address\x02Copy\x02Start\x02Service\x02Stop\x02" +
	"Local Directory\x02Add\x02Quick Add\x02Remote Desktop\x02HTTP File Serve" +
	"r\x02SOCKS5 Proxy\x02Enable\x02Open Config\x02Direct Edit\x02Copy Access" +
	" Address\x02Delete proxy \x22%[1]s\x22\x02Are you sure you would like to" +
	" delete proxy \x22%[1]s\x22?\x02Disable proxy \x22%[1]s\x22\x02Are you s" +
	"ure you would like to disable proxy \x22%[1]s\x22?\x02Disable\x02Port" +
	"\x02Passive Port Range\x02Error"

var es_ESIndex = []uint32{ // 170 elements
	// Entry 0 - 1F
	0x00000000, 0x00000015, 0x00000025, 0x00000039,
	0x00000056, 0x00000069, 0x00000084, 0x0000009c,
	0x000000ab, 0x000000c0, 0x000000d0, 0x000000d6,
	0x000000f0, 0x0000010c, 0x00000123, 0x0000016d,
	0x000001c6, 0x00000204, 0x00000234, 0x0000023c,
	0x00000245, 0x00000254, 0x00000269, 0x0000028f,
	0x00000296, 0x000002a6, 0x000002bc, 0x000002d3,
	0x000002ef, 0x00000318, 0x0000031f, 0x0000032c,
	// Entry 20 - 3F
	0x0000033d, 0x00000355, 0x0000038a, 0x000003b1,
	0x000003d1, 0x00000411, 0x00000427, 0x0000042f,
	0x00000436, 0x0000044e, 0x00000461, 0x00000469,
	0x0000046e, 0x00000476, 0x0000047e, 0x00000489,
	0x00000491, 0x0000049b, 0x000004af, 0x000004be,
	0x000004d3, 0x000004e8, 0x000004f1, 0x00000554,
	0x00000568, 0x00000588, 0x0000058e, 0x0000059d,
	0x000005a3, 0x000005ae, 0x000005b5, 0x000005bb,
	// Entry 40 - 5F
	0x000005c3, 0x00000625, 0x0000062d, 0x00000637,
	0x00000641, 0x0000064c, 0x0000065c, 0x0000067a,
	0x0000067c, 0x0000068b, 0x000006a0, 0x000006bb,
	0x000006c4, 0x000006cc, 0x000006e1, 0x000006ed,
	0x00000710, 0x00000725, 0x00000751, 0x00000761,
	0x00000785, 0x0000078e, 0x00000796, 0x000007a4,
	0x000007bc, 0x000007cb, 0x000007f9, 0x00000826,
	0x0000082e, 0x00000846, 0x0000088c, 0x000008a8,
	// Entry 60 - 7F
	0x000008d7, 0x000008e4, 0x000008ee, 0x000008f3,
	0x000008f8, 0x00000902, 0x00000910, 0x00000921,
	0x0000092e, 0x0000093c, 0x00000951, 0x00000962,
	0x00000976, 0x00000981, 0x00000999, 0x000009a2,
	0x000009ae, 0x000009be, 0x000009ca, 0x000009dc,
	0x000009e3, 0x000009eb, 0x000009f7, 0x00000a04,
	0x00000a15, 0x00000a29, 0x00000a32, 0x00000a39,
	0x00000a43, 0x00000a5e, 0x00000a69, 0x00000a9e,
	// Entry 80 - 9F
	0x00000aae, 0x00000ac2, 0x00000ac8, 0x00000ad7,
	0x00000ae8, 0x00000aed, 0x00000b01, 0x00000b14,
	0x00000b1e, 0x00000b4c, 0x00000b5f, 0x00000b85,
	0x00000b94, 0x00000ba0, 0x00000ba7, 0x00000bb0,
	0x00000bbb, 0x00000bc2, 0x00000bc9, 0x00000bdb,
	0x00000be2, 0x00000beb, 0x00000bf4, 0x00000bff,
	0x00000c10, 0x00000c18, 0x00000c28, 0x00000c3a,
	0x00000c54, 0x00000c61, 0x00000c6b, 0x00000c80,
	// Entry A0 - BF
	0x00000c91, 0x00000cad, 0x00000cc4, 0x00000cfb,
	0x00000d16, 0x00000d4f, 0x00000d5c, 0x00000d63,
	0x00000d7b, 0x00000d81,
} // Size: 704 bytes

const es_ESData string = "" + // Size: 3457 bytes
	"\x02Administrador de FRP\x02Versión: %[1]s\x02Versión FRP: %[1]s\x02Fech" +
	"a de compilación: %[1]s\x02Todos los archivos\x02Archivos de configuraci" +
	"ón\x02Archivos de certificado\x02Archivos clave\x02Archivos de registro" +
	"\x02Nueva versión!\x02Sobre\x02Descargar actualizaciones\x02Comprobando " +
	"actualizaciones\x02Buscar actualizaciones\x02Para comentarios o para inf" +
	"ormar errores, visite la página del proyecto:\x02Para ver la documentaci" +
	"ón de configuración de FRP, visite la página del proyecto FRP:\x02Se pr" +
	"odujo un error al buscar una actualización de software.\x02Actualmente n" +
	"o hay actualizaciones disponibles.\x02Aceptar\x02Cancelar\x02Configuraci" +
	"ón\x02Nueva Configuración\x02Importar configuración desde archivo\x02Ed" +
	"itar\x02Abrir documento\x02Mostrar en la carpeta\x02Importar desde archi" +
	"vo\x02Importar desde portapapeles\x02Exportar todas las configuraciones " +
	"a ZIP\x02Borrar\x02Nueva Config\x02Ajustes manuales\x02Importar configur" +
	"ación\x02Ya existe otra configuración con el nombre \x22%[1]s\x22.\x02No" +
	" se puede copiar el archivo \x22%[1]s\x22.\x02Eliminar configuración " +
	"\x22%[1]s\x22\x02¿Está seguro de que desea eliminar la configuración " +
	"\x22%[1]s\x22?\x02Editar Configuración\x02Básico\x02Nombre\x02Dirección " +
	"del servidor\x02Puerto de servicio\x02Usuario\x02Auth\x02Método\x02Ningu" +
	"na\x02Simbólico\x02Secreto\x02Audiencia\x02Dirección de token\x02Autenti" +
	"cación\x02Latidos del corazón\x02Conexión de trabajo\x02Registro\x02* Dé" +
	"jelo en blanco para no registrar ningún registro y eliminar el archivo d" +
	"e registro original.\x02Archivo de registro\x02Seleccionar archivo de re" +
	"gistro\x02Nivel\x02Días máximos\x02Admin\x02Dirección\x02Puerto\x02Clave" +
	"\x02Recurso\x02Seleccione un directorio local desde el que el servidor d" +
	"e administración cargará los recursos.\x02Depurar\x02Conexión\x02Protoco" +
	"lo\x02Proxy HTTP\x02Conectar cuenta\x02Tiempo de espera de conexión\x02s" +
	"\x02TCP Keep-Alive\x02Intervalo de latidos\x02Tiempo de espera de latido" +
	"\x02Encender\x02Apagado\x02Nombre de anfitrión\x02Certificado\x02Selecci" +
	"onar archivo de certificado\x02Clave de certificado\x02Seleccionar archi" +
	"vo de clave de certificado\x02CA de confianza\x02Seleccionar archivo CA " +
	"de confianza\x02Avanzado\x02Mux TCP\x02Mux Keepalive\x02Dirección de la " +
	"fuente\x02Otras opciones\x02Salir después de fallar el inicio de sesión" +
	"\x02Desactivar el inicio automático al arrancar\x02Disfraz\x02Opciones p" +
	"ersonalizadas\x02* Consulte la sección [common] del archivo de configura" +
	"ción de FRP.\x02La configuración ya existe\x02El nombre de configuración" +
	" \x22%[1]s\x22 ya existe.\x02Editar Proxy\x02Aleatorio\x02Tipo\x02Role" +
	"\x02Visitante\x02Llave secreta\x02Dirección local\x02Puerto local\x02Pue" +
	"rto remoto\x02Dirección de enlace\x02Puerto de enlace\x02Nombre del serv" +
	"idor\x02Subdominio\x02Dominios personalizados\x02Ruta URL\x02Multiplexor" +
	"\x02Usuario de ruta\x02Banda ancha\x02Versión de proxy\x02Vacío\x02Cifra" +
	"do\x02Compresión\x02Usuario HTTP\x02Contraseña HTTP\x02Reescritura de ho" +
	"st\x02Enchufar\x02Nombre\x02Ruta Unix\x02Seleccione la ruta de Unix\x02R" +
	"uta local\x02Seleccione una carpeta para la lista de directorios.\x02Pre" +
	"fijo de tira\x02Equilibrio de carga\x02Grupo\x02Clave de grupo\x02Cheque" +
	"o de salud\x02Tipo\x02Se acabó el tiempo\x02Recuento de fallas\x02Interv" +
	"alo\x02* Consulte los parámetros admitidos por FRP.\x02El proxy ya exist" +
	"e\x02El nombre de proxy \x22%[1]s\x22 ya existe.\x02Abrir registro\x02De" +
	"sconocido\x02Correr\x02Detenido\x02Comenzando\x02Parada\x02Estado\x02Dir" +
	"ección remota\x02Copiar\x02Comienzo\x02Servicio\x02Deténgase\x02Director" +
	"io local\x02Agregar\x02Añadir rápido\x02Escritorio remoto\x02Servidor de" +
	" archivos HTTP\x02Proxy SOCKS5\x02Habilitar\x02Abrir configuración\x02Ed" +
	"ición directa\x02Copiar dirección de acceso\x02Eliminar proxy \x22%[1]s" +
	"\x22\x02¿Está seguro de que desea eliminar el proxy \x22%[1]s\x22?\x02De" +
	"shabilitar proxy \x22%[1]s\x22\x02¿Está seguro de que desea desactivar e" +
	"l proxy \x22%[1]s\x22?\x02Deshabilitar\x02Puerto\x02Gama de puertos pasi" +
	"vos\x02Error"

var ja_JPIndex = []uint32{ // 170 elements
	// Entry 0 - 1F
	0x00000000, 0x00000014, 0x0000002c, 0x00000048,
	0x00000063, 0x0000007c, 0x0000008f, 0x000000a5,
	0x000000bb, 0x000000ce, 0x000000ea, 0x000000ee,
	0x0000010a, 0x00000126, 0x0000013c, 0x000001ac,
	0x0000021e, 0x00000273, 0x000002b3, 0x000002b6,
	0x000002c6, 0x000002cd, 0x000002dd, 0x0000030e,
	0x00000315, 0x0000032b, 0x00000341, 0x00000363,
	0x0000038e, 0x000003bc, 0x000003c3, 0x000003d3,
	// Entry 20 - 3F
	0x000003e0, 0x000003f9, 0x0000043b, 0x0000046f,
	0x00000488, 0x000004b9, 0x000004c9, 0x000004d0,
	0x000004d7, 0x000004f0, 0x00000503, 0x00000510,
	0x00000517, 0x00000524, 0x0000052b, 0x00000538,
	0x00000542, 0x0000054c, 0x0000055f, 0x00000566,
	0x00000576, 0x00000583, 0x0000058a, 0x000005f3,
	0x00000606, 0x00000622, 0x0000062c, 0x00000639,
	0x00000643, 0x00000659, 0x00000669, 0x00000679,
	// Entry 40 - 5F
	0x00000680, 0x000006e7, 0x000006f4, 0x000006fb,
	0x0000070b, 0x0000071d, 0x00000733, 0x0000074c,
	0x0000074e, 0x00000762, 0x00000778, 0x0000079a,
	0x000007a1, 0x000007a8, 0x000007b5, 0x000007bf,
	0x000007de, 0x000007ee, 0x0000081c, 0x0000082f,
	0x00000861, 0x00000868, 0x00000872, 0x0000088b,
	0x000008a1, 0x000008b7, 0x000008d6, 0x00000901,
	0x0000090e, 0x0000092a, 0x0000097d, 0x0000099f,
	// Entry 60 - 7F
	0x000009cd, 0x000009e3, 0x000009f0, 0x000009fa,
	0x00000a01, 0x00000a0e, 0x00000a18, 0x00000a31,
	0x00000a47, 0x00000a5d, 0x00000a76, 0x00000a8c,
	0x00000a9c, 0x00000aaf, 0x00000ac8, 0x00000adf,
	0x00000af5, 0x00000b0b, 0x00000b15, 0x00000b31,
	0x00000b35, 0x00000b3f, 0x00000b46, 0x00000b58,
	0x00000b6d, 0x00000b86, 0x00000b96, 0x00000ba9,
	0x00000bb5, 0x00000bca, 0x00000bdd, 0x00000c1d,
	// Entry 80 - 9F
	0x00000c3c, 0x00000c49, 0x00000c56, 0x00000c6c,
	0x00000c79, 0x00000c83, 0x00000c96, 0x00000ca0,
	0x00000cb3, 0x00000ced, 0x00000d15, 0x00000d49,
	0x00000d65, 0x00000d75, 0x00000d85, 0x00000d8c,
	0x00000d93, 0x00000d9a, 0x00000da1, 0x00000dba,
	0x00000dc4, 0x00000dce, 0x00000ddb, 0x00000de5,
	0x00000e04, 0x00000e0b, 0x00000e1e, 0x00000e3d,
	0x00000e5b, 0x00000e6f, 0x00000e76, 0x00000e86,
	// Entry A0 - BF
	0x00000e93, 0x00000eb8, 0x00000ee0, 0x00000f17,
	0x00000f3f, 0x00000f82, 0x00000f89, 0x00000f93,
	0x00000faf, 0x00000fb9,
} // Size: 704 bytes

const ja_JPData string = "" + // Size: 4025 bytes
	"\x02FRP マネージャ\x02バージョン：%[1]s\x02FRP バージョン：%[1]s\x02コンパイル日：%[1]s\x02すべてのフ" +
	"ァイル\x02設定ファイル\x02証明書ファイル\x02秘密鍵ファイル\x02ログファイル\x02新しいバージョン！\x02約\x02更新を" +
	"ダウンロード\x02アップデートの確認\x02更新を確認する\x02コメントやバグの報告については、プロジェクトページにアクセスしてください" +
	"：\x02FRP 設定ドキュメントについては、FRP プロジェクトページにアクセスしてください：\x02ソフトウェアアップデートの確認中にエ" +
	"ラーが発生しました。\x02現在、利用可能なアップデートはありません。\x02OK\x02キャンセル\x02設定\x02新しい設定\x02フ" +
	"ァイルから設定をインポートする\x02編集\x02ファイルを開く\x02フォルダで見て\x02ファイルからインポート\x02クリップボードか" +
	"らインポート\x02すべての設定をZIPにエクスポート\x02削除\x02新しい設定\x02手動設定\x02設定のインポート\x02\x22" +
	"%[1]s\x22 という名前の別の設定が既に存在します。\x02ファイル \x22%[1]s\x22 をコピーできません。\x02設定 " +
	"\x22%[1]s\x22 を削除\x02本当に設定 \x22%[1]s\x22 を削除しますか？\x02編集の設定\x02基本\x02名前" +
	"\x02サーバーアドレス\x02サーバポート\x02ユーザー\x02認証\x02認証方法\x02なし\x02トークン\x02秘密鍵\x02受信者" +
	"\x02トークンのURL\x02認証\x02接続を維持\x02作業接続\x02ログ\x02* ログを記録せず、元のログファイルを削除するには、空" +
	"白のままにします。\x02ログファイル\x02ログファイルを選択\x02レベル\x02最大日数\x02管理者\x02管理者アドレス\x02管" +
	"理ポート\x02パスワード\x02資産\x02管理サーバーがリソースをロードするローカルディレクトリを選択します。\x02デバッグ\x02接" +
	"続\x02プロトコル\x02HTTP プロキシ\x02接続プールの数\x02接続タイムアウト\x02s\x02TCP 接続を維持\x02接続" +
	"を維持間隔\x02接続を維持タイムアウト\x02有効\x02無効\x02ホスト名\x02証明書\x02証明書ファイルを選択\x02証明書キー" +
	"\x02証明書キーファイルを選択します\x02信頼できる CA\x02信頼できる CA ファイルを選択します\x02高度\x02多重化\x02多" +
	"重化接続を維持\x02送信元アドレス\x02別のオプション\x02ログイン失敗後に終了\x02起動時に自動起動を無効にする\x02カスタム" +
	"\x02カスタムオプション\x02* FRP 設定ファイルの [common] セクションを参照してください。\x02設定はすでに存在します" +
	"\x02設定名 \x22%[1]s\x22 は既に存在します。\x02プロキシの編集\x02ランダム\x02タイプ\x02役割\x02ビジター" +
	"\x02秘密鍵\x02ローカルアドレス\x02ローカルポート\x02リモートポート\x02バインドアドレス\x02バインドポート\x02サーバー" +
	"名\x02サブドメイン\x02カスタムドメイン\x02URL ルーティング\x02マルチプレクサ\x02ルートユーザー\x02帯域幅\x02" +
	"プロキシバージョン\x02空\x02暗号化\x02圧縮\x02HTTP ユーザー\x02HTTP パスワード\x02ホストの書き換え\x02" +
	"プラグイン\x02プラグイン名\x02Unix パス\x02Unix パスを選択\x02ローカルパス\x02ディレクトリリストのフォルダを選" +
	"択します。\x02プレフィックスを削除\x02負荷平衡\x02グループ\x02グループ秘密鍵\x02健康診断\x02タイプ\x02タイムアウ" +
	"ト\x02失敗数\x02チェック間隔\x02* FRP 対応のパラメータをご参照ください。\x02プロキシはすでに存在します\x02プロキシ" +
	"名 \x22%[1]s\x22 は既に存在します。\x02ログフォルダを開く\x02わからない\x02ランニング\x02停止\x02起動" +
	"\x02停止\x02状態\x02リモートアドレス\x02コピー\x02始める\x02サービス\x02止まる\x02ローカルディレクトリ\x02追" +
	"加\x02クイック追加\x02リモートデスクトップ\x02HTTP ファイルサーバー\x02SOCKS5 プロキシ\x02有効\x02設定を" +
	"開く\x02直接編集\x02アクセスアドレスのコピー\x02プロキシ \x22%[1]s\x22 を削除します\x02本当にプロキシ " +
	"\x22%[1]s\x22 を削除しますか？\x02プロキシ \x22%[1]s\x22 を無効にする\x02プロキシ \x22%[1]s" +
	"\x22 を無効にしてもよろしいですか？\x02無効\x02ポート\x02パッシブポート範囲\x02エラー"

var ko_KRIndex = []uint32{ // 170 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000e, 0x0000001c, 0x0000002e,
	0x00000043, 0x00000051, 0x0000005f, 0x00000070,
	0x0000007e, 0x0000008c, 0x0000009e, 0x000000a9,
	0x000000c3, 0x000000d7, 0x000000eb, 0x00000144,
	0x00000195, 0x000001e7, 0x0000021d, 0x00000224,
	0x0000022b, 0x00000232, 0x0000023d, 0x0000025e,
	0x0000026b, 0x00000279, 0x0000028a, 0x000002a4,
	0x000002c4, 0x000002ed, 0x000002f4, 0x000002ff,
	// Entry 20 - 3F
	0x0000030d, 0x00000321, 0x0000035d, 0x0000038b,
	0x000003a1, 0x000003cd, 0x000003db, 0x000003e8,
	0x000003ef, 0x000003fd, 0x0000040b, 0x00000415,
	0x0000041c, 0x0000042a, 0x00000431, 0x00000438,
	0x00000443, 0x00000451, 0x0000045c, 0x00000463,
	0x0000046e, 0x0000047c, 0x00000486, 0x000004e0,
	0x000004ee, 0x00000503, 0x0000050a, 0x00000518,
	0x00000522, 0x00000533, 0x00000541, 0x0000054e,
	// Entry 40 - 5F
	0x00000555, 0x000005a8, 0x000005b2, 0x000005b9,
	0x000005c0, 0x000005cf, 0x000005da, 0x000005ef,
	0x000005f1, 0x00000600, 0x00000615, 0x00000631,
	0x00000638, 0x0000063f, 0x00000650, 0x0000065a,
	0x00000672, 0x00000680, 0x0000069c, 0x000006b4,
	0x000006da, 0x000006e4, 0x000006ee, 0x00000703,
	0x00000711, 0x0000071f, 0x0000073b, 0x00000761,
	0x00000768, 0x00000780, 0x000007bb, 0x000007da,
	// Entry 60 - 7F
	0x00000811, 0x00000822, 0x0000082f, 0x00000836,
	0x0000083d, 0x00000847, 0x00000852, 0x00000860,
	0x0000086e, 0x0000087c, 0x0000088d, 0x0000089e,
	0x000008ac, 0x000008bd, 0x000008d8, 0x000008e6,
	0x000008f6, 0x00000907, 0x00000911, 0x00000922,
	0x00000930, 0x0000093a, 0x00000941, 0x00000950,
	0x00000962, 0x00000976, 0x00000983, 0x00000997,
	0x000009a3, 0x000009b6, 0x000009c4, 0x00000a00,
	// Entry 80 - 9F
	0x00000a14, 0x00000a22, 0x00000a29, 0x00000a3b,
	0x00000a49, 0x00000a50, 0x00000a5e, 0x00000a6c,
	0x00000a73, 0x00000ab1, 0x00000ad3, 0x00000b0d,
	0x00000b22, 0x00000b36, 0x00000b40, 0x00000b4a,
	0x00000b51, 0x00000b58, 0x00000b5f, 0x00000b6d,
	0x00000b74, 0x00000b7b, 0x00000b85, 0x00000b8c,
	0x00000ba0, 0x00000bad, 0x00000bbb, 0x00000bcf,
	0x00000be2, 0x00000bf3, 0x00000bfa, 0x00000c08,
	// Entry A0 - BF
	0x00000c16, 0x00000c2e, 0x00000c47, 0x00000c76,
	0x00000c95, 0x00000cca, 0x00000cd1, 0x00000cd8,
	0x00000cf0, 0x00000cf7,
} // Size: 704 bytes

const ko_KRData string = "" + // Size: 3319 bytes
	"\x02FRP 관리자\x02버전: %[1]s\x02FRP 버전: %[1]s\x02빌드 날짜: %[1]s\x02모든 파일\x02구성" +
	" 파일\x02인증서 파일\x02열쇠 파일\x02로그 파일\x02새로운 버전!\x02에 대한\x02업데이트 다운로드\x02업데이트 " +
	"확인\x02업데이트 확인\x02의견을 보거나 버그를 보고하려면 프로젝트 페이지를 방문하세요:\x02FRP 구성 문서를 보려면 " +
	"FRP 프로젝트 페이지를 방문하십시오:\x02소프트웨어 업데이트를 확인하는 동안 오류가 발생했습니다.\x02현재 사용 가능한 업데" +
	"이트가 없습니다.\x02확인\x02취소\x02구성\x02새 구성\x02파일에서 구성 가져오기\x02편집하다\x02파일 열기" +
	"\x02폴더에 표시\x02파일에서 가져오기\x02클립보드에서 가져오기\x02모든 구성을 ZIP 으로 내보내기\x02삭제\x02새 " +
	"구성\x02수동 설정\x02구성 가져오기\x02이름이 \x22%[1]s\x22 인 다른 구성이 이미 있습니다.\x02\x22%" +
	"[1]s\x22 파일을 복사할 수 없습니다.\x02\x22%[1]s\x22 구성 삭제\x02\x22%[1]s\x22 구성을 삭제하" +
	"시겠습니까?\x02구성 수정\x02기초적인\x02이름\x02서버 주소\x02서버 포트\x02사용자\x02인증\x02인증 방법" +
	"\x02없음\x02토큰\x02비밀 키\x02받는 사람\x02토큰 URL\x02입증\x02대기 중\x02작동 연결\x02통나무" +
	"\x02* 로그를 기록하지 않고 원본 로그 파일을 삭제하려면 비워 둡니다.\x02로그 파일\x02로그 파일 선택\x02수준\x02" +
	"최대 일수\x02관리자\x02관리자 주소\x02관리 포트\x02비밀번호\x02자산\x02관리 서버가 리소스를 로드할 로컬 디렉" +
	"토리를 선택하십시오.\x02디버그\x02연결\x02규약\x02HTTP 프록시\x02연결 수\x02연결 시간 초과\x02s" +
	"\x02TCP 대기 중\x02연결 간격 유지\x02연결 시간 초과 유지\x02켜다\x02폐쇄\x02호스트 이름\x02자격증\x02" +
	"인증서 파일 선택\x02인증서 키\x02인증서 키 파일 선택\x02신뢰할 수 있는 CA\x02신뢰할 수 있는 CA 파일 선택" +
	"\x02고급의\x02다중화\x02다중화 대기 중\x02소스 주소\x02다른 옵션\x02로그인 실패 후 종료\x02부팅 시 자동 시" +
	"작 비활성화\x02관습\x02사용자 지정 옵션\x02* FRP 설정 파일의 [common] 부분을 참고하세요.\x02구성이 이" +
	"미 있습니다.\x02구성 이름 \x22%[1]s\x22 이(가) 이미 존재합니다.\x02프록시 편집\x02무작위의\x02유형" +
	"\x02역할\x02방문객\x02비밀 키\x02지역 주소\x02로컬 포트\x02원격 포트\x02바인드 주소\x02바인드 포트\x02" +
	"서버 이름\x02하위 도메인\x02사용자 정의 도메인\x02URL 라우팅\x02멀티플렉서\x02경로 사용자\x02대역폭\x02" +
	"프록시 버전\x02비어 있는\x02암호화\x02압축\x02HTTP 사용자\x02HTTP 비밀번호\x02호스트 재작성\x02플러" +
	"그인\x02플러그인 이름\x02Unix 경로\x02선택 Unix 경로\x02로컬 경로\x02디렉토리 목록에 대한 폴더를 선택하" +
	"십시오.\x02스트립 접두사\x02부하 분산\x02그룹\x02그룹 비밀 키\x02건강 체크\x02유형\x02시간 초과\x02실" +
	"패 횟수\x02간격\x02* FRP 에서 지원하는 매개변수를 참조하십시오.\x02프록시가 이미 있습니다.\x02프록시 이름 " +
	"\x22%[1]s\x22 이(가) 이미 존재합니다.\x02로그 폴더 열기\x02알려지지 않은\x02달리기\x02중지됨\x02시작" +
	"\x02멎는\x02상태\x02원격 주소\x02복사\x02시작\x02서비스\x02중지\x02로컬 디렉토리\x02추가하다\x02빠른 " +
	"추가\x02원격 데스크탑\x02HTTP 파일 서버\x02SOCKS5 프록시\x02켜다\x02구성 열기\x02직접 편집\x02액" +
	"세스 주소 복사\x02프록시 \x22%[1]s\x22 삭제\x02\x22%[1]s\x22 프록시를 삭제하시겠습니까?\x02프록" +
	"시 \x22%[1]s\x22 비활성화\x02\x22%[1]s\x22 프록시를 비활성화하시겠습니까?\x02폐쇄\x02포트\x02" +
	"패시브 포트 범위\x02오류"

var zh_CNIndex = []uint32{ // 170 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000e, 0x0000001d, 0x00000030,
	0x00000045, 0x00000052, 0x0000005f, 0x0000006c,
	0x00000079, 0x00000086, 0x00000096, 0x0000009d,
	0x000000aa, 0x000000bd, 0x000000ca, 0x00000107,
	0x00000145, 0x00000164, 0x00000183, 0x0000018a,
	0x00000191, 0x00000198, 0x000001a5, 0x000001bb,
	0x000001c2, 0x000001cf, 0x000001e5, 0x000001f5,
	0x00000208, 0x0000022b, 0x00000232, 0x0000023f,
	// Entry 20 - 3F
	0x0000024c, 0x00000259, 0x00000289, 0x000002a7,
	0x000002bf, 0x000002fe, 0x0000030b, 0x00000312,
	0x00000319, 0x00000329, 0x00000339, 0x00000343,
	0x0000034a, 0x00000357, 0x0000035b, 0x00000362,
	0x00000369, 0x00000373, 0x00000380, 0x00000387,
	0x00000394, 0x000003a1, 0x000003a8, 0x000003e7,
	0x000003f4, 0x00000407, 0x0000040e, 0x0000041b,
	0x00000422, 0x0000042f, 0x0000043c, 0x00000443,
	// Entry 40 - 5F
	0x00000450, 0x00000484, 0x0000048b, 0x00000492,
	0x00000499, 0x000004a5, 0x000004b5, 0x000004c2,
	0x000004c6, 0x000004d7, 0x000004e4, 0x000004f1,
	0x000004f8, 0x000004ff, 0x0000050c, 0x00000519,
	0x0000052c, 0x00000539, 0x00000552, 0x00000562,
	0x0000057b, 0x00000582, 0x0000058f, 0x0000059f,
	0x000005af, 0x000005bc, 0x000005d8, 0x000005ee,
	0x000005f8, 0x00000608, 0x00000638, 0x00000648,
	// Entry 60 - 7F
	0x00000669, 0x00000676, 0x00000683, 0x0000068a,
	0x00000691, 0x0000069b, 0x000006a2, 0x000006af,
	0x000006bc, 0x000006c9, 0x000006d6, 0x000006e3,
	0x000006f0, 0x000006fa, 0x0000070a, 0x00000715,
	0x0000071f, 0x0000072c, 0x00000739, 0x00000746,
	0x0000074a, 0x00000757, 0x00000764, 0x00000770,
	0x0000077c, 0x00000788, 0x0000078f, 0x0000079c,
	0x000007a8, 0x000007bb, 0x000007c8, 0x000007f6,
	// Entry 80 - 9F
	0x00000803, 0x00000810, 0x0000081d, 0x0000082a,
	0x00000837, 0x00000844, 0x00000851, 0x0000085e,
	0x0000086b, 0x0000088b, 0x0000089b, 0x000008bc,
	0x000008d2, 0x000008d9, 0x000008e6, 0x000008f0,
	0x000008fd, 0x0000090a, 0x00000911, 0x0000091e,
	0x00000925, 0x0000092c, 0x00000933, 0x0000093a,
	0x00000947, 0x0000094e, 0x0000095b, 0x00000968,
	0x0000097a, 0x00000988, 0x0000098f, 0x000009a2,
	// Entry A0 - BF
	0x000009af, 0x000009c2, 0x000009da, 0x00000a01,
	0x00000a19, 0x00000a40, 0x00000a47, 0x00000a4e,
	0x00000a61, 0x00000a68,
} // Size: 704 bytes

const zh_CNData string = "" + // Size: 2664 bytes
	"\x02FRP 管理器\x02版本：%[1]s\x02FRP 版本：%[1]s\x02构建日期：%[1]s\x02所有文件\x02配置文件" +
	"\x02证书文件\x02密钥文件\x02日志文件\x02发现更新！\x02关于\x02下载更新\x02正在检查更新\x02检查更新\x02如有任" +
	"何意见或报告错误，请访问项目地址：\x02了解 FRP 软件配置文档，请访问 FRP 项目地址：\x02检查更新时出现错误。\x02当前没有" +
	"可用的更新。\x02确定\x02取消\x02配置\x02新建配置\x02从文件导入配置\x02编辑\x02打开文件\x02在文件夹中显示" +
	"\x02从文件导入\x02从剪贴板导入\x02导出所有配置 (ZIP 压缩包)\x02删除\x02新建配置\x02手动设置\x02导入配置" +
	"\x02另一个同名的配置「%[1]s」已存在。\x02无法复制文件 \x22%[1]s\x22。\x02删除配置「%[1]s」\x02确定要删除" +
	"配置「%[1]s」吗？此操作无法撤销。\x02编辑配置\x02基本\x02名称\x02服务器地址\x02服务器端口\x02用户名\x02认证" +
	"\x02认证方式\x02无\x02令牌\x02密钥\x02接收者\x02令牌地址\x02鉴权\x02心跳消息\x02工作连接\x02日志\x02" +
	"* 留空则不记录日志，且删除原来的日志文件。\x02日志文件\x02选择日志文件\x02级别\x02最大天数\x02管理\x02管理地址\x02" +
	"管理端口\x02密码\x02静态资源\x02选择管理服务器使用的静态资源目录。\x02调试\x02连接\x02协议\x02HTTP 代理" +
	"\x02连接池数量\x02连接超时\x02秒\x02TCP 保活周期\x02心跳间隔\x02心跳超时\x02开启\x02关闭\x02主机名称" +
	"\x02证书文件\x02选择证书文件\x02密钥文件\x02选择证书密钥文件\x02受信任证书\x02选择受信任的证书\x02高级\x02多路复" +
	"用\x02复用器心跳\x02使用源地址\x02其他选项\x02初次登录失败后退出\x02禁用开机自启动\x02自定义\x02自定义参数" +
	"\x02* 参考 FRP 配置文件的 [common] 部分。\x02配置已存在\x02配置名「%[1]s」已存在。\x02编辑代理\x02随机" +
	"名称\x02类型\x02角色\x02访问者\x02私钥\x02本地地址\x02本地端口\x02远程端口\x02绑定地址\x02绑定端口" +
	"\x02服务名称\x02子域名\x02自定义域名\x02URL 路由\x02复用器\x02路由用户\x02带宽限流\x02代理版本\x02空" +
	"\x02加密传输\x02压缩传输\x02HTTP 用户\x02HTTP 密码\x02Host 替换\x02插件\x02插件名称\x02Unix " +
	"路径\x02选择 Unix 路径\x02本地路径\x02选择需要显示目录列表的文件夹。\x02移除前缀\x02负载均衡\x02分组名称" +
	"\x02分组密钥\x02健康检查\x02检查类型\x02检查超时\x02错误次数\x02检查周期\x02* 参考 FRP 支持的参数。\x02代" +
	"理已存在\x02代理名「%[1]s」已存在。\x02打开日志文件夹\x02未知\x02正在运行\x02已停止\x02正在启动\x02正在停止" +
	"\x02状态\x02远程地址\x02复制\x02启动\x02服务\x02停止\x02本地目录\x02添加\x02快速添加\x02远程桌面\x02" +
	"HTTP 文件服务\x02SOCKS5 代理\x02启用\x02打开配置文件\x02直接编辑\x02复制访问地址\x02删除代理「%[1]s」" +
	"\x02确定要删除代理「%[1]s」吗？\x02禁用代理「%[1]s」\x02确定要禁用代理「%[1]s」吗？\x02禁用\x02端口\x02被" +
	"动端口范围\x02错误"

var zh_TWIndex = []uint32{ // 170 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000e, 0x0000001d, 0x00000030,
	0x00000045, 0x00000052, 0x0000005f, 0x0000006c,
	0x00000079, 0x00000086, 0x00000096, 0x0000009d,
	0x000000aa, 0x000000bd, 0x000000ca, 0x00000107,
	0x00000145, 0x00000164, 0x00000183, 0x0000018a,
	0x00000191, 0x00000198, 0x000001a5, 0x000001bb,
	0x000001c2, 0x000001cf, 0x000001e5, 0x000001f5,
	0x00000208, 0x0000022b, 0x00000232, 0x0000023f,
	// Entry 20 - 3F
	0x0000024c, 0x00000259, 0x00000289, 0x000002a7,
	0x000002bf, 0x000002fe, 0x0000030b, 0x00000312,
	0x00000319, 0x00000329, 0x00000339, 0x00000343,
	0x0000034a, 0x00000357, 0x0000035b, 0x00000362,
	0x00000369, 0x00000373, 0x00000380, 0x00000387,
	0x00000394, 0x000003a1, 0x000003a8, 0x000003e7,
	0x000003f4, 0x00000407, 0x0000040e, 0x0000041b,
	0x00000422, 0x0000042f, 0x0000043c, 0x00000443,
	// Entry 40 - 5F
	0x00000450, 0x00000484, 0x0000048b, 0x00000492,
	0x00000499, 0x000004a5, 0x000004b5, 0x000004c2,
	0x000004c6, 0x000004d7, 0x000004e4, 0x000004f1,
	0x000004f8, 0x000004ff, 0x0000050c, 0x00000519,
	0x0000052c, 0x00000539, 0x00000552, 0x00000562,
	0x0000057b, 0x00000582, 0x0000058f, 0x0000059f,
	0x000005af, 0x000005bc, 0x000005d8, 0x000005ee,
	0x000005f8, 0x00000608, 0x00000638, 0x00000648,
	// Entry 60 - 7F
	0x00000669, 0x00000676, 0x00000683, 0x0000068a,
	0x00000691, 0x0000069b, 0x000006a2, 0x000006af,
	0x000006bc, 0x000006c9, 0x000006d6, 0x000006e3,
	0x000006f0, 0x000006fa, 0x0000070a, 0x00000715,
	0x0000071f, 0x0000072c, 0x00000739, 0x00000746,
	0x0000074a, 0x00000757, 0x00000764, 0x00000770,
	0x0000077c, 0x00000788, 0x0000078f, 0x0000079c,
	0x000007a8, 0x000007bb, 0x000007c8, 0x000007f6,
	// Entry 80 - 9F
	0x00000803, 0x00000810, 0x0000081d, 0x0000082a,
	0x00000837, 0x00000844, 0x00000851, 0x0000085e,
	0x0000086b, 0x0000088b, 0x0000089b, 0x000008bc,
	0x000008d2, 0x000008d9, 0x000008e6, 0x000008f0,
	0x000008fd, 0x0000090a, 0x00000911, 0x0000091e,
	0x00000925, 0x0000092c, 0x00000933, 0x0000093a,
	0x00000947, 0x0000094e, 0x0000095b, 0x00000968,
	0x0000097a, 0x00000988, 0x0000098f, 0x000009a2,
	// Entry A0 - BF
	0x000009af, 0x000009c2, 0x000009da, 0x00000a01,
	0x00000a19, 0x00000a40, 0x00000a47, 0x00000a4e,
	0x00000a61, 0x00000a68,
} // Size: 704 bytes

const zh_TWData string = "" + // Size: 2664 bytes
	"\x02FRP 管理器\x02版本：%[1]s\x02FRP 版本：%[1]s\x02構建日期：%[1]s\x02所有文件\x02配置文件" +
	"\x02證書文件\x02密鑰文件\x02日誌文件\x02發現更新！\x02關於\x02下載更新\x02正在檢查更新\x02檢查更新\x02如有任" +
	"何意見或報告錯誤，請訪問項目地址：\x02了解 FRP 軟件配置文檔，請訪問 FRP 項目地址：\x02檢查更新時出現錯誤。\x02當前沒有" +
	"可用的更新。\x02確定\x02取消\x02配置\x02新建配置\x02從文件導入配置\x02編輯\x02打開文件\x02在文件夾中顯示" +
	"\x02從文件導入\x02從剪貼板導入\x02導出所有配置 (ZIP 壓縮包)\x02刪除\x02新建配置\x02手動設置\x02導入配置" +
	"\x02另一個同名的配置「%[1]s」已存在。\x02無法複製文件 \x22%[1]s\x22。\x02刪除配置「%[1]s」\x02確定要刪除" +
	"配置「%[1]s」嗎？此操作無法撤銷。\x02編輯配置\x02基本\x02名稱\x02服務器地址\x02服務器端口\x02用戶名\x02認證" +
	"\x02認證方式\x02無\x02令牌\x02密鑰\x02接收者\x02令牌地址\x02鑑權\x02心跳消息\x02工作連接\x02日誌\x02" +
	"* 留空則不記錄日誌，且刪除原來的日誌文件。\x02日誌文件\x02選擇日誌文件\x02級別\x02最大天數\x02管理\x02管理地址\x02" +
	"管理端口\x02密碼\x02靜態資源\x02選擇管理服務器使用的靜態資源目錄。\x02調試\x02連接\x02協議\x02HTTP 代理" +
	"\x02連接池數量\x02連接超時\x02秒\x02TCP 保活週期\x02心跳間隔\x02心跳超時\x02開啟\x02關閉\x02主機名稱" +
	"\x02證書文件\x02選擇證書文件\x02密鑰文件\x02選擇證書密鑰文件\x02受信任證書\x02選擇受信任的證書\x02高級\x02多路復" +
	"用\x02復用器心跳\x02使用源地址\x02其他選項\x02初次登錄失敗後退出\x02禁用開機自啟動\x02自定義\x02自定義參數" +
	"\x02* 參考 FRP 配置文件的 [common] 部分。\x02配置已存在\x02配置名「%[1]s」已存在。\x02編輯代理\x02隨機" +
	"名稱\x02類型\x02角色\x02訪問者\x02私鑰\x02本地地址\x02本地端口\x02遠程端口\x02綁定地址\x02綁定端口" +
	"\x02服務名稱\x02子域名\x02自定義域名\x02URL 路由\x02復用器\x02路由用戶\x02帶寬限流\x02代理版本\x02空" +
	"\x02加密傳輸\x02壓縮傳輸\x02HTTP 用戶\x02HTTP 密碼\x02Host 替換\x02插件\x02插件名稱\x02Unix " +
	"路徑\x02選擇 Unix 路徑\x02本地路徑\x02選擇需要顯示目錄列表的文件夾。\x02移除前綴\x02負載均衡\x02分組名稱" +
	"\x02分組密鑰\x02健康檢查\x02檢查類型\x02檢查超時\x02錯誤次數\x02檢查週期\x02* 參考 FRP 支持的參數。\x02代" +
	"理已存在\x02代理名「%[1]s」已存在。\x02打開日誌文件夾\x02未知\x02正在運行\x02已停止\x02正在啟動\x02正在停止" +
	"\x02狀態\x02遠程地址\x02複製\x02啟動\x02服務\x02停止\x02本地目錄\x02添加\x02快速添加\x02遠程桌面\x02" +
	"HTTP 文件服務\x02SOCKS5 代理\x02啟用\x02打開配置文件\x02直接編輯\x02複製訪問地址\x02刪除代理「%[1]s」" +
	"\x02確定要刪除代理「%[1]s」嗎？\x02禁用代理「%[1]s」\x02確定要禁用代理「%[1]s」嗎？\x02禁用\x02端口\x02被" +
	"動端口範圍\x02錯誤"

	// Total table size 23038 bytes (22KiB); checksum: 80129B53