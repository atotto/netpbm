// Test PPM files

package netpbm

import (
	"bytes"
	"compress/flate"
	"image"
	"testing"
)

// ppmRaw is a complete, raw PPM file, compressed with Flate compression.
const ppmRaw = "Ԛ\tL\xdbG\xd6\xc0\xa5|\x95\x12E_\xa5]\xb5\xd2vUm\x97\xddF]U[Uj\xb6\xcd\xeeF\xedn\xa26i\xb2m\x1aڒ&\xb4iK\xb8\x12RN\x1b\x8c9\x8c\x8d\x0fl\xb0\xb1\xf1}\x1f\x18c\x1b\x83\xef\x13\f6\xd8@8\xcai\xcedI/\x8e\x14\x04\xe40\xe1\n\xb0Ch\xb3Y\x02]\xa7\xf6\x9fdGO\xc8\xd8\xf3\x1f\xfff\xe6\xbd7\xef\xbdq\xf8\x81\xbdχ\xbd\xf9\xc1\x91\xc3\x11'?8\x18\xf6\xf6\xf1\x13\xe1a\xe1\xef\x9d\b;\x9a\x90\x98\x1e\a\v\xfb(\x0e\x06OHI\x0e{\xe5\xe5W\xf6\x1ex5\xec\xc0\xab{\xf7\xbf\xf6\xda\u07b5е\x95\x95\x95\xc5\xc5ť\xa5\xc5\xee\xae.\x9d\xa6\\\xc8$K\xd9$1\x83\xa0U\b]\xb5\xd57o\xdeX^^\xde\xe8v{\xce_\xefr\x9bM\xe6\xd9\xe9\x99;++k\x8f\xba\xad\xae\xad\x8d\x8d\x8d\xb546\x88Yd\x16\x1eVAG\xe8\x05h~Q\x16\r\x9f\x9d\x87H*\xc1\xc1u\x82|\x12\"ڠ\x14\x0e\r\xf4\xdbMF6>M\xc3̒S`\x89\x9f\u007f\xa8.W\xcc\xcd\xdd~\x84\xf0\xe3\xe3\xd7zڼrV\x01@\x1at\n\xbc\x1a\n\x8f\x92K\xa3R\xffy\xf5\xab\x8d\x0e\x13\xdfO\x1a\f\x86R\x0eQ\xcf\xcd\xe1\x11\xd2\xf2\x92?\xd5rP*N\xbeIL\x18v\x89\xa5Ep\x8bA\xbbtwkv\xbe\xb9kk\xedj\xbeKI\xf19\xf8\xfd5BW9\t\x8bL\xf2\xf9|\x0f\xf6\xbc\xe5\xf7\xf3\xa8\xf8V}Iiqƫ/\xed;|\xe8\xd0\xfe\x97_4\n\xf3\xfb\xab\xf92\x0e\xf9ƍ\x1b;L~\xcb\u007f[]*\xb2\xca\n\xbf4\xb1z\xec\xbcn\x1bo\xc0)\xc2\xc2Ϲ]u\x9bz\xce\xcd\xcd\xc1\xe1\xf0\xf0\xf0pr!\x89\x88\x88\x06\xddʨ\x99\u007fx\xfe\xb7\xbf\xfcœ\x16\t~\xa8V\xa4\x14Rw\x98\xbf\xbf\xaf_HE\xb7\xe8\x19\x00\xfb\x9e\xb4\xea\xe9*1crjjSg$\x12\xb9gϞ\xa7\x9ez\xeaȑ\xa3t\x02\xb2\xd7\xc1_\xefl`\x9e|\xeb\xafV)\x01l\x99ݨYXX\xdc\x19\xf2\x9b7o\xe5\xe5\xe6\x11\xe0QCu\xe2\xfb\xe1{\xec|\xab\x18kі\x03\a\xb4\xe9\x11\x18\f\xb6{\xf7\xee}\xfb\xf6ed \xcaE\x8c\x065\x19t\x06\x8f\xb4\x99X\x99\t\xa7yD\xd8\xe5\x01\xdf\x0e\x90\xdfYY\x1d\x1c\x18L\x8c\xfd\x8c\x95\x9fP\xa7.1\xc9\xc8V9\xc5!/l30\xc1\x92\x02\xa4\x1ay\x81\xb6\\\xe4\xf7\xcfmzptt4&&\xe6رc*\xb5ڦS\xb8\x95\x85\x1b\xfc\xe0/ز<X|wW7\xe4\xf0wV\x9a\x1b\xbd\x94\xbc\x94\x12t\xa2N%5\x18\x8c\x1eo\x93\xb3\xcem5\xebMJ\xaeQ\x90\x0f\x8c\xb7\xc7\xc6\x15\x95`\x81#\xddn\x90\xd9\xd9\xd9RVA\x97\x95so\xd7|\xd5\x02\x87\x8c\xc0 \x13!u\xa1+\xab\xab\x8d\r\r2j\x16\"\xe1c\xbb\xcd\x02\xfe\xbd\xffSpd9LU\x0e)\x16\x18\xa6\x82\x9e\xd5\xd4ذ\xdd8\xad\x97\x1a\xcbh\x99\x80\xf9~\xc5\x03\x13\x97\x92\xd3\x1d6\x1bt\xfcc\xa3\xa3l\x02\\FE\xbajk6}\xa4\xd1h\"\"\"h4Z.,~\xb0V\x04\x1c\x119\xf7\x8b\xcbW\xae<8\xc8\xe0\xe0\x10%/\xb1\xdbƽ\x1f\xfe\xc7]\xe0c\xd2\x13\xa0\xe3G!Қ*)&\xb5\xe8\xe6-\xff\xfd\xef{<\x1e`\x95\xbbv\xed:x\xf0`fz\x9a\x96\x8b\x02k\xdbW#`\xe1\xe1R\xa9\xf4\xfa\xf5\xeb?\xea\xff\x98B&\x92P\x80\xe7\x11<\b\xbf\xce\xef\x10\x18\x05\xe8\xd2R\x05\x14\xf0\xce\x1aW!\"\xaa\xb6\x8c\xa8SI6i\x8e\xcb\xe5z\xee\xb9\xe7\x9ex\xe2\t\x14\nUY\xa1fa\x93\x01\x89\xb7\x92V_Q\xacff\xc7G\x1eO\x8e\x8bĤǣS\xcevXؾm\xe07d\xd8%\x89\xfd䃕\xd5P\x92_\xbb\xf6=\xf0l\xa7\xde?\xd9nd\x9a\x84\x18\xa5\x8c\xbfi\xf8\x85\x85\x05\x1c\x0e\a\x1cKvv\xf6\xcc\xf4\x145/\xb1(+\xf6/\xfb\xff\x18\xf6\xbb\u07ff\xf5Ɵ\xeaUd\xbb\x94\xd0XI\x05v\xf1\x13\xe4\xf7\f\x99GH\xd6\xe9L\xa1\x82\xef\xea\xec\x161HUܼ\xb3'\x0f\x0fԊj\x15D\x93F\xb6\xb0\xb8\xb4]\xff\xd1o\xbf>\xfa\xc6\xfe\xbd{v\xef\xfd\xff'w\xef\xd9\xf3\xb7\x03/\xe1\xe1\x9f\xe5%\u007f\x9c\x9b\x14\xd9XE\xfb\xaf\xfc@\x80v\xa5%'\x85\xe6x\xed\x1fB$\xc6z5\xc5W=r\xe4\x85\xd3\xfdNQ\x97\x85S%)\x1e\x1f\xdf\xd67r\xe8\xc5\xf1g\x8e\xfd߮]\xbf\xf9\xf5\xd3g\xc3\x0f+h\x19\x06~\x9eY\x84\xc9M<\xa3\xe5\xe6\x05\xc2\x0f\x8e\x83bt\xea\xc4\xc4T\x90\xf0\xcb\xcbw\xa4b\t*\xf1\xccHC)\xf0l'\xde\xfc\xf3P\x9d\x04\x9cPF!\xa6T\xcc_\xd9JGU\xca\xf2\xe2\x9c\xf3\x88\xf3\xa7D\xa4\x14Y\x11L\xc3\xca2\n\xd1\x06>J\xcd\xc8\x04V`\x95\xe0\x03\xe1\a\xe2VSe\xd2\xd2 \xf9\x81VSHDy1\xbc\xae\xbc\xe8\xd9g\x9e>\xf2\xfa+\x80\x1f\f\xdeW-\u0530s\xf0\xe8,\xbf\xff\xdf^h~~\x9eH\xc0\xd1P\x17\f\x02\fP\x18=\x1f\xa5\xe3\xe6rqI\x91'\xdfDf\xa4E}rJ\xc5\xc8\xea\xfdI˽_@@%\xa1\x17\\\xbb6\x15\xdc\xfa/k\xcaˤ\x85iW\xeae\xa7O\x1c\xda\xffҾa\x97\xf4\x9e\x95\xf5\xday\xe8\x94\xcfȸ\\\x1a)_X\x82\xa7a\x12[t%`j@I\x90\x17>R3\x91\x05\x19QU\x95\x956\x9b\x1d\x8c\xd3\xe7\xf3\xa99\xb8n+7@\xfe^;_\xcbE\xf3\xd9,\xa0\x03\xc1L\xa1\xd1\xdb\x04\x8f=\x05B\xfao\x9b\xcb3\xe2O\r8\xff#T\x03\xfed\xa0F\xd0_#\xe8\xab\x16\xac\x87\rw\x83\x19\xb0t\x85\xc8\x18\x12\xe2\\L\xe4\xbb\xcbK\x8b\xab\xab\xab\x0e\x87\x03\fe\xa8\x90ի\x8b\x03\xe4\a\xd2af\xe7$}\xee\xaas\a\xc3?==C\xc0\xe20\xa9\xe7@JE\xca<\xe7\v@\x01\xc0,6\xa28\x8b\b}\xa9\xb9\t\f\xd2\xeb\xf3)U\x15\x15r^\x93\x96\x1e8\xff\xba\x962\xb3`\xc9I \xd0\rf\nSSӽ==n\xa7\xfd\xd33\x1f6ki\x1b\x8b\x1c\x88\xf4W\v0\x99I\xaezO\x95R\f\xf2\xdc&\x1d\xb33`\xfd\xb9\xb7\xbf\x1f\xbf\xf7\x168}BRR\x98\x9c\x9c\x12\x10S\xc1\xb2\x04\xf8\xed \x04\x02\xb9\xc9%\x03\xa7\xc3\xc2\xeb\xb4pz\x1e\x86\xfcǈN$,H\u0380\xc1\xe5\xf22\xbd^\xafRU|\xf5\xd57\xc1\xccBH\xc7\xf5U\xf3\xbb\x1f\x9e\xe4g\v\b\xc5iE\x84\x1f\x8c\xb1\xb1\xd1l\xb6\xae\x04Qf1\xe9\xb5Vq\xbe/`7\x18\xbc\f\u05c9\x93\xa3#\xc6\xc7\xc7A\x10\xd8\xd4\xd4\xe4v\aeγ3\xb38D\x02\xf09;\xc6\x0fL\xa00\xf3\\\\\xecy\xa7\xd3\xd9\xd0\xe0\x99\x9e\x9e\x0e\x86\u007fii\xb9\xc6jp\xca\v\x02\xb7\xe2 \x05\x98\x9b\x8a\x8e\x8c\x8b\xbd\x10\xaap\xee\xea\xc8\x15\x9d\x88\xd4W\x1d\x9a-\xb8\xebl\x05\x1b\x85\x88\xed:\xb4\x19\x99/\xbc\xf0b\xa8\xf8A~Z)\xe79\xcbB\xb0\x05`\x84K:\xbaY\x94o\x93\xe2\xc0\x14z\xb7\x19p\xb0N\xf2\xec3\xbf\xf2\xfbC\x96\x17[,\xb6r\x16z\xcb\x04\xf0\xa1d\xd0)$\xa23\xec\x0e\x87\xd9l\xe60\xa9z!~\xcb)\fՉ\x8f\x1fzm\xe4\xea7!\xcch$<\xa6[M\t\x92\xbf\xd7\xc6Q\xcbE?\x945VW\xadZE\xad\x82\xf4\xe0\xb6\x02\x13\x8e\x8b\xfc\x87\xd7\xdb\x1cB\xfeɩ\x19\x061\xbb\xcb\x1a\xd4\x16(K\x10\xa5\"ng{K\xff\xc0\xd0\xec\xf5\x1bE\x05\x18{)\xf1A\xfe\xbe\x1a!\x16\x1e%\x93\x95\x856\x1d\xf6x\x9a\x10\xe7O_vK\u007f\x1e|\x053\x87\x92\x13_\xc5\xc7\x1aE\x04)\x03\xa7\x94qKKr[\r\x8c-\xd31\x05\r\x89\xc6\xe0B\x9eѧ\xa5\xc0\xd9\u0604@\xb2\xdaM\xd2e\xe5y5\xb4K:Ɔ\xd9\x02\x0f\x03\xb4\x11\x04\x9c[\xfa\x04\xf0\xa6[E)( CP\x02\xf5\x1f\xfd\xfb\xeb\x83\x0f\xcf\xff\xb0\x02fJ$\x92\xa1(\xaa8l6\x16\xe6\"\b\xfe!\f\x81\xec<O%M\xa9TC\xc1\x0f\u0092\xe2\xfc\x8c~(#\n\xc0\xdfXU\"\x91ʡ\xe0\xf7\xcf\xcdi\xe4B\xaf\x86\x02]D\x01\xf8\x81\xb1\x10\b\x85Д\xa3\xef\x18\xb5\x15\x02Bb\xa8\"\x8a-\xa5Y\xcb\xc8\xc7\x12 *\x8a\x1a\r\xc6\xf4\xd8\b\xe8L`\xc3\xff\xd0hL\x88\xf8[Z\xdbPiqm\x06&D\xfc\xc0\xc7\xeax\x18\x99\\\x05ٵ\xe9\x04.7\xdd.\xc5Cd\x02 \x84\xa6\xe4\x9c\xef\xec\x82\xeaji\xee\xf6B\x11\x01\xa3\xe3\xa2 \xe2\a\x9a\x89\xbc\x189::\x01ݽ\x00\x8fI\xad\xe2\xe4B\xc4?T'\t\u007f\xfb\x8dť;\xd0\xf1W($F\x01\x06\"\xfe\xc1Zq\xf8\xbb\xef@z\xa9g3hl\xb2\x82\x8d\xcb\xeb\x90;\u007f\x10\x17\xc5\xc6\xc4C\xcao7U\xd9\x01?$\xca/b\xe7_T\x96\xab \xe5\xb7\x1a*mRH\xd6\x1f(\xff\x17\x9f\x9e\x18\x19\xf9\x1aR~1\x9f\xad\x85\xc6\xff\xf4\xdayi\x17c::;\a\x06\x06fff \xe2\x17q\xe9P\xf8O`\xb9\\\\b˥\xf5\"0\xe0oh\xf0\x04Y\xcbݮi\x95R\x8b\x18\x1bZ\xfd\x01iWS\x15\xf5\xfd\xe3\x87FFF\xee\xe6\x1a7i4Zww\x0f\x14\xfc\xaej\xb3SQ\x14Z\xfb\xed\xb2pP\xa9\xd1v{\x8d\xc1`\xe2px\xf3\xf3\xf3t:sff\x16\n~\x8f\xc7[\xa3,\t!\xfc\x97&6=?͠\xff\xe1\x16\xb5\xb9\xb9\x05\xe4/##W!\xd2\xff\x91\xab_3I9\xedFV\b\xcaqv~CE1\xaf0\xc7bq\xecد\x98ܮ\x06rQ\xd1\xfa\x11\x10\x84\t\x838ǫ)V\xb3\xd14b~wO\xff\xda\x0e\xb6\xe1\xe1+\xce:w^fj\xab\x9e\xfe3\xa6\x00\x0e\xa9^\aOA\xcb,!d\xebu\xda\xdb\xf3\x8bk;ۖ\x96\x96\xa6\xa7\xa7GG\xc7pY\xa9\xbdvN\x80\x17\x04`\xa6\x03\xb5\xe2\xe1:1\x17\xfbEvJ\xbc\xddb\x1a\x1b\x1d_{\xa4\r\x1c1ٰ\xc4\x1a9a\xc0)Z\xff\tJ\xf5ֵe\xf0&\xf0\xed\xadz\x06*)\xf2\xec\xa9\xf7\r:\xfd\xc2\xfc<HE\xd7\x1e\x836==\x93\x9e\x92\x8a\x83G\xf3\x99E\x99\xf0d5\v\xc4ռ.+\xb7\xd3\x02\x84\xd3a洛\xd8\n\x1a\xf2\xdcG\xefD\x9e>\xd3\xdcܶ\xf6\xf8\xb5\x95\x95\xd5\nM\xd5\xe4\xe4\xfa\xedy}\x83'9:®\xa0\n\xc9\xd9\xec\x02D!\x06A*\xc0K\xa5\n\xbf\u007f~\xedqm\xc3×\xdbۿ\xdcx\xed\xf56Ri\x8c\xaen_wo\xdfwߍ\xaf=\xf6\xcd\xef\xf7\xdbl\x8e\x89\x89\x89\xbb\xa9\xf1\x18Du3\xe8\xda\xe4\xe4$\x93\xc9joo\a\xaf\xb1X\xfc\xe2\xe2\xe2\xff\x16?p#\x1d\x1d\x9d\x04\x02!**Z\xaf7\xef\xd8\xf7\xfe+\x00\x00\xff\xff"

// ppmPlain is a complete, plain PPM file, compressed with Flate compression.
const ppmPlain = "\xec][r\x1c\xb9\xb1\xfd\xc7**b\xfe'\xf0\x06j\x05\xf7\xf7~\xdc\rh\xaei\x8f\xc2\x1a9\xc2#\xc7l\xdfyN\x02U(\x8aM5ɦ\xd8d\xc3&\xa6\x9b\xadj\xb0\x1e\a\x89̓\xaf\xff\r\xe6\x97\xe5\xff~\xff\xfc\xe7\xf2\xf7\xcf_\xee\x96\xdf?\xfd\xb9|Z\xfe\xf8ϗo\x9f\xbf|\xfez\xb7\xfc\xff\xbf\xfe\xf8\xe3\xee\xeb\xb7_\u007f\xfd\xd5\xc8q\xf2\xb2\xfc\xf5\xf9\xdb\xef\xcbo_>}\xfd\xe7\x82#\xfe\\>\u007f]~\xbb\xfb\xf6\xd7\xdd\xddW9&\x87%'SJ1\xfc\x8f\x8cex\xbd\xff\xfe\xa1\xdfO}\xc6\xcf\xcdr\xc6\xffN|y\x9b\xe4G\a<v\x06O\x19\xaf0\x87\xdbGZ\xf7\x91\xbd\x0ey\x9fC\x94\xe18R\xb0Kr\xd5$\x1b\x96X\xfc\x12c]\xa2\x8fK\xac2\x92\f'\xa3\xc8g1-Ѧ%\xf9\xb2$kyl*\xab\x8c\xb0\xa4\x1c\x97\x12\xe4<\x82\xfc}\x99\xbb䴏\x93\xe7f\x97\x12\xe5\xb3(߯I\xe7\xaa\xde$\xf9;)\xca\xdc\xf2Z|\x95!\xe7\xec1W\x1d\xbe\xff\x1e\x9f˜\xe3\xe5s\xc4m\x8e\xbc\xd6}X\xc1\xce*\xb8\xacU\xb1\x9a\x9cb\x18\xaf\xd6/Ap\x96\x9c\xe0Wp\x15\xb25)\b\xd6\xf3\xba\x84Z\xe45˱V>_\x97\x18\xb2\x1e/\xd8\x0e\x82\u007f9v\tA\xd6D\x90cW\x19\x82Y\xac\x17\xf9\xae\tk\xe0\xda\xc01^\xfe-\xc9w\x93\xcd\\\x179\xc99%y\x8d\xf2\xb9\xbc'\x9ee\x0e\xae'\x9eO\xe4\x9a\xc3Z\xe3\xdf]\xe5sYW8\xcf\xe0\xd7\xc5W\x19Iօ\x93\xefɺ\xc4zӵRt\xddL||\xec9\xac\xc30\xc5\n\xceE\xa6f'\xd8\x02F֎\xe9\n\xfc,\x1d?9A~\x12\x97<\x8e\xff\xe6\x81\xed`\x88+\xac\t`\x1b\xb8\x97\x91\xe5\xd8\\\xb2\x8c\xf5\x91s\x109\xee\x04k.\x18b\xbb\x8d\xb8\xca\xdfO2\xb7\xc3ށu$k\xa4$\xfc-\xee\x0f\xc06\xd7NH\x82\xe3\xc0\xfdB\x8e7\xa1\xca\xf1v\xe5w\xb1v\xb0&\xb8V\xe5;9;\xca\u007f\xfe\x8e}\a\xeb\x11kM\xae\x8f\u007fW搽\xcb<\xbc\xa7L\x8c\xbd\xd79d\xbfW\xfd\xa1\b\xc6 'E\xdeA\x1eS\x17\x80,\x95瞡\x1f\b>\xf09\xe5i.\xbb\xeea\x9b\f^\xb3\xac\x83l \xabq|\x1f\x17\xbb\x96h\xa9\xa7`\x10\xe7\xc0\xbd\xc8\xf1\f\xf9\xcc5\xb5\xea\xba\b\xc0m1N\xaeŉ\x1e\xe4d\xfd8ٟܚtȞ\xe0d.'\xfa\f\xf7\x03\xec3kh\xd7\x15u/\xc2\x1e&s`\xff\xa2.\xd5\xc7\xc4\xd8\xfb\x9a\xa3c\x06r\f\xb8\x06>DN\xa6\xe4\r\x9f\xb1\x13l{\xafr09\xcay\xac\x01\xe0\v\xfayJ*S\xa1GP\xb6\xcbw\x89\xf3\x04}#\x1b];п\xf3ˮ%\xee\x03s\xf7\xe1E\x8e\xf7A]\xca\x15\xb5\x05D\xeew\x9dG\xce\xcd\xf4s\x8b\xa2\u007fa]\x049g\xe7\xd6\xc5\xc9ޅ\xad\bW.\x8b\x85\xb2<V\xd5Ǡ\x1fa\xdd\xf3~\xd8lhw\xc8w_\x80\xf7\x89\xd37\x98\x03r+G\x95\xcb\xfd\x19B\xa7\xa6>\x0e\xb9\x18\x939`\xf9\x9c\xf7]\xa6Bg\xb6\xd4\t\flO\xd1Cdd\x1d/\xb8\x16\xd9'\xb6\xe1c\xfb\x01\x1e\xf9\xe3Eb\xeb\xd03\xa0FÝ\xc7b\x172܅\xbe;\xdbLM\xccG\xb7xA\xbd+\xf8w\xaf\x9a\x8cM\x87\x9dJV\xb8\x81䇆VRnc\x9d(\xbb\xa69B\x1b\xf2̩ɊV\x1aڳT\xe9\x1b\x9a\x04\xcb\xd4X\a<\x9c\x89\xf6.\xe9\xb3\xe2\x04\xf3\x8bd\xa7\xb6\xeb\v\x99\x98\x03c\xf3\x12\xb4\xc7}x\xd7~\xa8\x96\xb4\x11\xa0\x8ec\x88HW\n(\x91\xea\x11S\xdap\xdb\xf2\xf1p\xc2G\xf0\xe7m\xfb\x82\xd8\x0f\xa2\xb6\xfb\xdcL\x01l\x13\xb1\x98,j\x8d\x9a\xc2a\x82\xec\xad\xe7\xd8(\xb4h`2\xf2\xb9\xc8+(\x87`a\xaee}\x96\xab>K\x00\x9c\x18\x80\xc0\xc2s\xaeJ\r\x12\xb4k\xa4iI\xd5|U\xf5\x80xh°\t̆\x99\x01\xf0P\xdfq\xbc\xdd\xc1\x0e\xccu\xf5\x9a\xaa\xc8\v\xeeG\x165\xbb\x8f\x00\xaa\xa5\r/*I\x1f\xd8VH\xad`\xc0t\x0eրb\xa1J\x05\xb3ÕӘ\x1f?\x1f\xb6*\x11\x10\x86\v\u0379\x81^\x9dj̛\xcc\x01\x9c7\xacg+\xa6%\xa8\xc1N\xb1\x91\x9a\xa8J\xe9U\xa5\xdbH\x81CՐ\u007f۩\x0eP)\x8e\xdf\xcb\xd6R\xcd\xe5z\x00F\x92\xd2\xd3x\xf6\xa4\xf6\x80\x19\xc8O\x98t\xa0\x157\x93NMU\xc5L28\x86\xf2\x95tb>P\x98\xaf~OI\xa5s\x18\xa5p\x94F\xa2\v\x00\xeb9\x95\xf3\xd43\x9a\xe3\xb2\xf6A5a}W\xa5p\xe8\x1e\xf0j\x1e\x9c\xa0*'N/>\ae9G\xa3\xbf\r)p\f\x9ab\xabR\xd8ؗ\xa9\x94V\x95w\xb1\x1e\xbe\xfb\xe4\xf3\x18\\J\xc0\x00\xd7\x00\xf6\x05\xe2\x81t\x8c\x19\xcdV\xba\x8bH\xbd\xafXK\x1bU\xff\xda\xf7T\xe4\xb3\xc96\xf3\xef\xa9y\x9a\xa9\xa74\fs\xad\x9e\x81\xf9\xa3R7\xack\xa5*\xe5z\x93RT\xbc\xb7)M\x9c\xbe\xe2\x1c*/\xb9\xbf\x1bu\xa7\xa8R\xbec>\x0fnɗ\x9d\a\x8c\xb4>\xbc\x18j}(\x95\xaeC\x8c\x83}\xdf'\r\xaf\xfa\xccH\xe7\x10#P\xf2\xbd\xc8|\xe8\xc3r^9_\ue7aa.'\xa6v\xaa\xfb\x1e\xe7Gl\xdb\xe5L3ܜ)\xff\xa9\xff\x84\xd5\x12\xfb\x90/\x13\xa7\xaf7G\xa8\xceP\x8eǨ\xf2\xd3;>g\xca]\xe8\xe7П\xc3@9>\xf1<T\xcf\xd11چ\x99\xffw\xa5lC\xa4\xa7\xe1\x13\x87\x03\x93\x0e&\x98\xc8\xee(Qc\xd3\xf2\xe1\xd8\xec\x0eX\xae\xd0؇)\xa2ul\xe3\xe4J\xef\x9a|\xd5\x1d\x84\xceݤ$\xa2X\x0e\x03-s\xae$\u007f&\xdau\xfeM\xbb\x83\xd3\x18\x1a\x14\xee\x98h\x86\x13\xa9/\x9f\x83\xdau.\xea\x04t\xc5(\xc1\x1cT\x9a\xae\xdd\xd9صt\xf7\x98\xf6\xf2c\xb4sw\xd0!\x8fP\u007fd\x15\xe9\x0f\xd0܇\xac3\x83\xb0\x01\xae1\xdařn\xcd\x03*\xc8q;\xee\x05\fcA\x98Ah\x94?ܳ\xc5\x1b\xec\r^\xce\xd8{\xcf\xef\aѻ\xf1\xfb&?\xc1\x95\xc0\x16\x11\xbb\xc4WYw^u4\xeco\xedo\x99g\"\xfcih\x1f\xdfw\x97,\xee\x05\xb8\x01\xb9s>\x01\xed\uec7b>\xd1~p\xa1\xb7\x01Y:\xf0`^tg\x0fB.\x8a\xe4t+\ta`E]➸\xa0MY\x95g8\xe3~\x9f>\x8f\x8dd\xce\x1a\xc2\xd2\xc6ȇ0\xdc\x058--\xd4\n!W.7:_u\x89\x1e\"\x03,8\xc1\x00]\x99\"\xbd\xe1ݴ\xf2\xab\x8d\xca}\xd7l\x1c\xd4 ,\x05,V\xaff\xc7ƅC\xf5\xe7r9z\xb6^\x04\xd5s\xd4\xf6\xf3M[ݸ4\x9a\xc1`\xb9R\xf0`D?\xe1\xfe\xd0F\r\xaa\xcb+\x0f\xde\xe9c\xfa:@\ad\xe5\xd0!\xd6\x00w\xb7\xd6f\x1a6\xef_\xd1\xe8&\nx\xaa3j\xben~\x91\x87\xf9\xef\x8b^\v\x95\x8c*\xe2\xb5\x06\xfa}\x04Äw\xb0\xf4\x0eA\xfe\xc9\xf2\x80\x0f\x13\xf49\xfdEN\x0ew9\xc1\xc9o\xba~\x03\xf6t\v|)M\xe3\x0f\x9e\x92^a\x1f\xb9K<\x00?\xf3B\xa8\xbfl\xc9\xec:\x94a\xf0\x0ev\xa8\xe0\xa7tos\xa8\x93s\xa5\xcbG\xd9\x14\x95\x82[\xa0U\xd0`\xaa\xb4*\x13N\xe7\xa2+\xf4\x84\xc8\xf7\xcc.\xe2\\\xdb\xd5\x1b\x16\xaa\x06\x8c\xb4`\xaf\xc7X\xc0\x8b]\v\x03\xb7\xb0\xcb\x00\xd2\xd0Jr\xe5\x0e\xb0R\xd3Ǻ\x03\xab\xddVkd\xec\nW'\f]\x9cv\xc8\xf4\x80\x02\xe6 @5&\xc16\xe3v\xe5t\xbbY\x10/'\x99/\xb6dv\xf2^n\x81\xa1\xd2\x15\x9dƺ1\xde-ߠS\xa8\xc7jW\xc3]\xbah\x9c\x06\x95\xe0\xf6L3\xa44\x94\x91\xd5\x1dH/\x0fQ'\xf7\xcd\xe1\x18\xf5\x88\x9aHspw\xfe\xecqN\xb6\xc5\xc0t\xa5fT\xe3/\x14Of5n,\xd9`\bq\xe0\xb3\xc5\xe8\xe2\x15\x8a\x17w(ל\x00th\xd9\x16\x1f\\\x9b#s\xed\xe7L|СE̻\xa6\xd4'uv\x0eJ\xbdz\xfb[l\xcb5h2'\xd6\v\x9d\x17N\xe34x\xedx\x96\xab\xbb=ў\xb2\xe9J1\xe3d[\\\x13\x1d\x93\xa1Ş\x96\xac\x0eq\xc1\x87ܶ\x12(\x18\xc1\x8c\x88\x9a\xe0t?0\f\x01\xa8\xee\xde~\x9aU\x8d,M\x9dIQ\x03d\xa1'@\xe0\b\xc0vog8\xefbR\xd8\x06O\x14\xac\xe7\xaa\x01\xebd\x87\\\a\xea\xaa\vm\xd0I4H\xb7\x01\x98\xc1\xeb*\xb0\xf9\nk@u\x14\xa3\xe1+m\xffBD\x00\x82\xbf\xc4\x06W\xcbO\xbd\xa6\n\xfe\xf5\x14\xef\xf2\x86\x82\xfdĢ\xdb\x03\x91t\x812\xa4\xa7\x0e\x96Y\xbe\r=fH\xc8!\xa3m\x1b\xafR\xd5sOf\x11\xfcKM\x14\f\xd0] ܷ\xa0U\xb9\xa7d\x047V \x9eeF\x11'\x8d\x05\x17\xec\x1be8\xb2n \xc0^\x0fLg\x10oK@\x82\xf0\xb5^\xb1\xbd\xaa\xbe\xa5\xcf\x11\x8a\xc8j\x0e\xebkP6v\xbd\xcb\xef\xec\v6zy\xde\xc03\xf9$\x1f\x95\xd5\xf1\xbb\xb7\x06\xe77\xde\x1f\rhk\xd7\xe1|\vX\xef8\xf7\x9dM\xbf*\xc1\xfe\x9d>ߣ-\xb6\xfb\xd3\x12]\x92&\x10\f\x9b\xfeM\x98\xa8=\x02\x00\xec9\x03\\\xc1h!\x11n\f.\x87\t\xc8d\x1bp\xdc\xea\xd1\xe11\x1bMQ\x9f\xfe\\F\x8a\x03\xef\x1b/\xbey\x117\xbe<\x1e?ר\x01\xf3`\x04\xca(\xd3\\\vh\xc4zْ\x91\xf4w*'\xa1'M\xf8\x16f\x18w\x8f1\xd6r\x1b\x87{\x05&ө\x1cPF\xd3S\x99z3\xfb\xf4\xa9\xeb\x85\xeci\xda\xf6]\xeeo\xb5s\xbe\x1a8}\x13\xec#\x13'\x95\xd1S}\xa13\x80\x96F\x0fprHB#~\x02?\x1f\x19\xee7\x93c\xa7\xa2\xaeNE\x96E\xb9\x16<o\x06 g\xaee\xac\xf7c\xc0\xb2\xdde;\r\x85uQ\x86?\xa9\xf2\x87k\xa1\x87M\xed\xe4\xab\xd0a\xce\xe2p\xee\xe9\x9c\xfaL\xd53\xdd\xf9\x9c!\xac\xf5&\f\xdc\x1e,\xee\xbd>w\x91\xa7\xc4KU\x99\f\xdc A\x8d\x1cH\n*\xfb}i\x1cݩ\xe8\xbf+\xdb\xf7\xb7ġ\xb6\xae\xa1g\x85\x96\x18\x02\xd9\r\x1f%t\x1e\x9b[b잀ʨ6\xe8T\x88\xd6\xf4u\xd4m\xde^\x87y\xea\xfb!\x11\x05ܫ\x87\x9ei\x8f\x1e\x98ۦ\xea\x87\xc4{&Rh2\x05]\x1d\x8c\xd4v\a\x19r\x15\xbc\xf4I\x8c\r\x89RC\xfa\x05#\xbd\xf0\x1e\x91\x8c݃J}\xa0\xeaސzT\xef;\xd0a\xce^\xfbE\t\xad\xa2v-e}j!\xff\x88\xb9\xe8I\xc3\xd0\u007f\x11\xadڇ]M\xdf\xfb\x90\x04\xccQo\x83\xfc\xd4\"\x17^m\xbaM\xf7\xce\xd7\xc7]\\\x9b]yu\xd7R\xb6\xe86r\b\xdeR\xcf=\x85\x0f\xb5\x01ۈj#`\xcf;\xc8Ə\x1a_\x135S\x81\xc5[\x90\x80ʔ\xa3\xf0T=\xff\x061ve\xd7B\xfeMu9\xfa\x99\x11\xc9\x01\"\xbb\xf9\x10Y4b\x88\xad\x83-\xd3\a\xf5\xde6nB\x0fb\xb6C2x\xa5\xdd\xdb\"\xe1\x0f<\xcc\xc4\xd8\xf5_\xcb\xc0\xa1)\xdf\x15\x1a\xc7U\rx<8\xf64;\xa7\r\xf2V\x95\xb6\xcf\u0381\xf8\x1b\xe1\xf9[\xca(x\xfe\\4S\t|^\xb1[\xd6\xdd#\xb6\xed\x94\xc9\xd7r-#\x9fC\xbf\x9b7\xca\xf3\xd6E1?cs\x1e\fEcQ\xa3\xa0\xbe\xac`7g\xff\x81c\x9f\x18;ů\x9ax?\xa6\xfe\xc0\xb5\xbe\xf6\xb5lr\xc9h&\xb1F\xc5:y\xcf\xe1\xd3\xc4\xfa\x83\x85f\xf6lK\x86\xd6*\x8f\xddd~\xe92\xdf<\"\xf3oH&\x0f\x18#_\xa4\x99\x89{0[\xd0ckl\xfc\xc0\x80\xff\xc1\xaf\xfa*\xd7\xc2\xec1\x8d'р\x93\xb2W\x9d\x9a\xa9\xb4\xa7\f۪\x037\x0fA\r\x89\xa9\x14\xe6\x19\xa9o\x1fS\xb0ke@\xad\xb0\xd4\xd3L\x92.\x82\xcd\x01\xdc\xc9\x10:M\x8bV\x80mD\xa1\x12\x8bt\xa0\\^\x80\f5T\x10E\xc7`\x139\x1fn\xe0\xd8\xc8\x1f&$g\xc2\xd5HfVG\x9c2}\xa4G\xf7\xf9r\xdb\xca6\x05\xb5bY+\xafi^\xc2\xe3eQZPD\xf2\xad,\x89\x1dr\x13⏌\xa6\xf3\xafe\xac\xe6\xd6\xd2\xc0\xbco\xc1\xe4\fP\t\x13\xeb?\x9a\xa3W\x02\xf3=\xd7\xc41\x85i+e\xf1\xb8\xcc\xff\x90\xc6)\x1d\xbcAI\xae^i\x90\xf7\x01y9N\x03\xe84\x89A\xab\xef\x91\x1c#I\x882>V\x9dGkSp\xf0}瞂y\xf3cc\xb6\xb4(\xf0\xbc\xd5Њ\xbdT\xc5:\xb1\xfe\x9ctm\x06\xb3\xf5\xc8v\x04.\xe6\xc8=`\b\xce\xfcpd\x8eV\xf1\xae\x1aP\xcd\xcaƉ\x81z\x9ak\x96\xb5\x8eZhQ\xfeV\x15y\x06-\xaf\xabV\x9a\xac\x1a\xf8\xac\xd5Z5\x00uw\"\xa7\xcb]\xcb`Do\x194\f\xc2F0\xa958\x9f\x89\xf5\xe7͡IjU\x1dܩU\xff\xf5\xd6<\x92O\xfa\x0eu\xfe\xbe\x97\x85!\xd1,/\xbbњN\x1b\xa7J\x9e\x98\xeft\xc2!0\xf2L\xd9n\x9e\x9b]\xa5\xa5\xd4V\x16j\xc0\xde\xc2\xe1\xe3\xc4\xfas\xe7\u0602\xf7\x98\xa3j\xdcꙈ\xb2U~\u007f:\xe6'\xf9y\xe9k\x91\xfd\xa5\xafM&\xe98\xf7\x1c\xccO\xac?\x14\xbc1\x94\xd7\xd2R\x1cp\xe6\xde/\xb37\x1dR?\xd7Y\x10Z\x80i+\u007fb\xdbp~b\xfd\x12s\f\x81\x1c\xccJDY\xd6T\x8e\xc9#S&\xff\xa4kكN\xf7b\x0eZ\x1eG\xb3h\xe3\xc4\xfaE:m\x84͡E\x9dwmN\xf0\xa1\xab\xcd\xc1\x893e\xf2+^ˈ\xf9\xdd\x06\xa1]m5\x18gb\xfd\x02s\xf4rŨ\x9e\x81\xb2þu+Cp\xbe\xd5\xf2\xdc9\x8any~\xc9ҩ\xc3\\\xe2ZzBPT\xc70x\x86\a\x02\x8c'\xd6_8\a1\xdfJn\xb3\x1c\x9eM\x06A\xadg\xea\xf6S\x87\xb9H \xfdw\xbc\xcd^\xe66։\xd3K\x971\x86?ơ\x83Xfg<r\x9a\b\xd4\x19[(L\x99\xfcz\xd7r\xa8Pd\xb5@\xc5ڊ\xccBױ\xeb\xc4\xe9\xa5\xe7`GK\xad\xb6Ջ\u007f\xb0t\xfaV\x1fҟ\xc3KO\x1d\xe6\xb9\xd72\x04\x15ﾳ\b\x1f1\xfd\r/h\xd13\xb1\xfeXw=\xc4\xee\xa5=A\x8d~\xf6^\x88\xa9\xba\xf7\x93l\xfe.\x03\xe9w\x1d\x92\xfeW\x1f\fbnzqqV\xfd\x9b8}\xb59zl\x1a\xe3\tّ[\xbb\x11?\x80\xf9\xa9\xc3\\\xeaZz\\eEQ7T\xb0ͬi\xbb\xf7\x94\xab\x13\xa7\xaf8\a\vʰa\x81\x1bZB\xd9\xc9[\xbej\x01\x90d\xb4\xa6tծ\xb5(\xfa\xe4g_͟5\x87\xb6\xd1\xcbZ\\0\xa5\xd6\x16lL<\x99:̥}\xaf\x87Bg\xb6\xb5;tn\xe2\xf4g\xf1\x96\xa9\xee\xad\xcfh\xc3V\xb3'\xd8N\x99|\xb9x\xfe\xa1\xd0M\x8a\xad=i/\xe0\xc6b6\x13\xa7?i\x8e-\xa1\x1cET[\xa2\xdc1\xa9|\xea0\x97I\x06\xdctx\xb6\x8dc\xa1:\xb4\x1ebnT`\xaf\x8a\x89ӟ\x94cU41\x1a\xb2\xe7\xd0\xf4\xeci9AS\x879\xbbp\xb16\x98\x83\x8c\x87N\x03=\x87y,\x18\xcc_ѡ\xc5]\x92\xb6,\x9a8\xbdl\xa3\x90\x96?\xc4\xe4\xe8\xf3\xb8ʩÜ\x9d\xf7\xdaۄ\x0eE\xf8[\x91s-\\{\xafh\xad\xe2]\xf3\x1b\xc6\xe26O/\xe8:\xb1\xfe\xc8\x1c\x8c\xbba۵\xa8\xc5#Y$-\xffH\xb7\x99:̩k\x19\v\x89\xb0\xa8r$\x0f\xcc\xd6w-\xdf{\xd3gʽ\xbcװ\xe7<\x8c\xedv\xb5\xb5\xbe\xe6\xafL\xac_\xa00~\xd1\x02\xc2\xcc\xffG\xf1,\xf0\x96ޞ\xb2a\xa7\x0e\xf3\xc3kٛ|i\\\x87c\x91>m\xbd^\xb5(W\xd0\x18\x0f\xb6{L\x85\x85i\xf1\xddC\x93\x97\xe8ۨ˦\xe3\xf8:\xb1~\x819\xf6v\xc7Q\x9fG˝\x9e:\xcc\x13\xae\x05]\xaf\x88\xdf؊鶦\xd0A\x1b\x0e\x9d蕩\r\xf5lhc=\xe4\x8c\xec{A\x9e8\xbd\xf0\x1c{M\x95\xe6\x8fu\xda\x1c\x03\xb5\x01\x98+\x9a܍\xeb0\x87\x06\x13f\x8f\x87\xac[Kg\xfa\xb2\x9d\xc6\xcdlu\x87\xb2\x9b\x18\xbb\xd6\xdaP\xb57B\xd0\xdcc\xc6\xdf\xf8\xf6lw\xff\xd4\xf2\x8c\xbaF\x1f\"\x1e\xa6\xe5t\x98\xad]9\xb8\xf5\xda\xf4m\xd4TYæk\xe3{\x13c\xd7?\a\nD\xeb`S}6(\xc4:\xd0\x06iA\x1b=8\xdf\x1aH}\xb0x\xcb\xc1O\xb4۞a\xa8\xbf\x81\xae\xa0-\xa7\xa6\xb4\x82\xe8A\x9b\xfdh\x8c^hM\xb2&\xc6\xde\xd5\x1cZ\xf4X\x8b\x1f\x0f\x05\x9f{\x8c\x1f\x1a+k\x01\xf0\xb55:i<3si\x0f\x05\xf0\xafL\x87\x19\xd6\xe6\xd63\xbd\x1e\x8a\xc0+\u007f\x98\xe9#b\x1b\u007f\xd4EA=\x1c4'\x8d\xea\x9f\xc6\xf5k\xac\xa3\x9b\x18\xbb\xa596\xde@\xe5\x1d\x9b\xc2\xf8&#W4\xaeO\xc7\x06\x85\x87Ɲ\xf9\xd8\xc4\xf3\xe1\"\xc0\x17h:\x9f\xcdȇ\xef{SR\xfe\xa45\xd0gGr\x91\xcd\xe0\x0e\xd5ƴ<N\xbb\xf8\x06C\xd9=\xe4QN|\xcc9\x86\x06\x88\x86y\x85\x90\x8d\xd1\xe9^\xc0\x1at\xab6E\x8d\xad\v8\xf4#4\a\r\xa9\xe9\xc3n\x8cU6\x1b\xbf\xd7k\fjc\xa9\xe3\xfb4\xd6Y*j_7\x9eD\xf4\x10\xd3뀡\x81\x01\x1a\xd5\"F\x91\xf8F\xc3]tQ\xef\xdcȭ4&\x9es\xbc\xe2\x1c{\x87}mrY\xe9\x83\x19\x1b\xc0\x1d\xed\xe5\xd4}\x8bfl.Cޚ:t\xab\a\x8c\xe6ͥq\u007f\xa8\x1b\t?\xb1+\xad\x01\xa1\xd6K\x95\xcf\r\x9b8#>\xa8\x12\xfbC-\xeb0\x9f\xed\x9c\xe3U\xe6\xd8\xedߨuO1N7\u007f\x01\ab\xb4\xf6nQ|ou \x95\x0fe\xb3X\xbeZm\f\xbez\x8e{\xba\xf4|.s\x8e\x9f;G\x97\xa51\xec\x8d\xec\xb3\xe6\xd6n\xe3\x98{\xcb\xfc[\u0602Ϩ\xd35\x9f˜\xe3\xed\xe7\xd8\xf4\x96\xe6â\xbe\xedv9\xef\x0eܝц\xe7\x99y\xe5\xe4\xf5\xfa\x98\xf7t\xce\xf1\x0ekCi\x1c\xac\xc6\xc2n5\xb8,\xf4\xf4`PO\x9dq&\xbd\xc1oL\xf3\x9e\xce9\xe6\x1cs\x8ek\x98\xc3\xfc\xb2\xfcϧ\u007f\xff\xf6\xe9\x1fw˧o\xcb\xdd\u05ff\x99\xff\x06\x00\x00\xff\xff"

// Determine if image.DecodeConfig can decode the configuration of a raw PPM
// file.
func TestDecodeRawPPMConfig(t *testing.T) {
	r := flate.NewReader(bytes.NewBufferString(ppmRaw))
	defer r.Close()
	cfg, str, err := image.DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	if str != "ppm" {
		t.Fatalf("Expected \"ppm\" but received %q", str)
	}
	if cfg.Width != 64 || cfg.Height != 64 {
		t.Fatalf("Expected a 64x64 image but received %dx%d", cfg.Width, cfg.Height)
	}
}

// Determine if image.Decode can decode a raw PPM file.
func TestDecodeRawPPM(t *testing.T) {
	r := flate.NewReader(bytes.NewBufferString(ppmRaw))
	defer r.Close()
	img, str, err := image.Decode(r)
	if err != nil {
		t.Fatal(err)
	}
	if str != "ppm" {
		t.Fatalf("Expected ppm but received %s", str)
	}
	_, ok := img.(Image)
	if !ok {
		t.Fatal("Image is not a Netpbm image")
	}
}

// Determine if netpbm.DecodeConfig can decode the configuration of a raw PPM
// file.
func TestNetpbmDecodeRawPPMConfig(t *testing.T) {
	r := flate.NewReader(bytes.NewBufferString(ppmRaw))
	defer r.Close()
	cfg, err := DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Width != 64 || cfg.Height != 64 {
		t.Fatalf("Expected a 64x64 image but received %dx%d", cfg.Width, cfg.Height)
	}
}

// Determine if netpbm.Decode can decode a raw PPM file.
func TestNetpbmDecodeRawPPM(t *testing.T) {
	r := flate.NewReader(bytes.NewBufferString(ppmRaw))
	defer r.Close()
	img, err := Decode(r, nil)
	if err != nil {
		t.Fatal(err)
	}
	if img.Format() != PPM {
		t.Fatalf("Expected PPM but received %s", img.Format())
	}
}

// Determine if netpbm.Decode can decode a raw PPM file with non-default
// options.
func TestNetpbmDecodeRawPPMOpts(t *testing.T) {
	r := flate.NewReader(bytes.NewBufferString(ppmRaw))
	defer r.Close()
	img, err := Decode(r, &DecodeOptions{
		Target: PPM,
		Exact:  true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if img.Format() != PPM {
		t.Fatalf("Expected PPM but received %s", img.Format())
	}
}

// Determine if image.DecodeConfig can decode the configuration of a plain PPM
// file.
func TestDecodePlainPPMConfig(t *testing.T) {
	r := flate.NewReader(bytes.NewBufferString(ppmPlain))
	defer r.Close()
	cfg, str, err := image.DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	if str != "ppm" {
		t.Fatalf("Expected \"ppm\" but received %q", str)
	}
	if cfg.Width != 63 || cfg.Height != 65 {
		t.Fatalf("Expected a 63x65 image but received %dx%d", cfg.Width, cfg.Height)
	}
}

// Determine if image.Decode can decode a plain PPM file.
func TestDecodePlainPPM(t *testing.T) {
	r := flate.NewReader(bytes.NewBufferString(ppmPlain))
	defer r.Close()
	img, str, err := image.Decode(r)
	if err != nil {
		t.Fatal(err)
	}
	if str != "ppm" {
		t.Fatalf("Expected ppm but received %s", str)
	}
	_, ok := img.(Image)
	if !ok {
		t.Fatal("Image is not a Netpbm image")
	}
}

// Determine if netpbm.DecodeConfig can decode the configuration of a plain PPM
// file.
func TestNetpbmDecodePlainPPMConfig(t *testing.T) {
	r := flate.NewReader(bytes.NewBufferString(ppmPlain))
	defer r.Close()
	cfg, err := DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Width != 63 || cfg.Height != 65 {
		t.Fatalf("Expected a 63x65 image but received %dx%d", cfg.Width, cfg.Height)
	}
}

// Determine if netpbm.Decode can decode a plain PPM file.
func TestNetpbmDecodePlainPPM(t *testing.T) {
	r := flate.NewReader(bytes.NewBufferString(ppmPlain))
	defer r.Close()
	img, err := Decode(r, nil)
	if err != nil {
		t.Fatal(err)
	}
	if img.Format() != PPM {
		t.Fatalf("Expected PPM but received %s", img.Format())
	}
}

// Determine if netpbm.Decode can decode a plain PPM file with non-default
// options.
func TestNetpbmDecodePlainPPMOpts(t *testing.T) {
	r := flate.NewReader(bytes.NewBufferString(ppmPlain))
	defer r.Close()
	img, err := Decode(r, &DecodeOptions{
		Target: PPM,
		Exact:  true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if img.Format() != PPM {
		t.Fatalf("Expected PPM but received %s", img.Format())
	}
}
