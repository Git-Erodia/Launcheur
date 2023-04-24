package Minecraft

import "os"

type variable struct {
}

var Username string = os.Getenv("username")
var HomeDir string = `C:\Users\` + Username + "\\"
var DekstopDir string = HomeDir + "\\Desktop\\"
var AppDataRoaming string = HomeDir + "\\AppData\\Roaming\\"
var AppDataLocal string = HomeDir + "\\AppData\\Local\\"
var AppDataLocalLow string = HomeDir + "\\AppData\\LocalLow\\"
var TempDir string = AppDataLocal + "Temp\\"
