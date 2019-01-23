let state;
let pki;
let parser = new DOMParser();
const charset = "MNBVCXY123ASDFGHJKL654QWERTZUIOP789qwertzuioplkjhgfdsamnbvcxy";

function setup() {
    pki = forge.pki;
    state = new State();
    let data = state.loadState();
    if (data != false && data != null) {
        aplyState(data);
    } else {
        newState();
    }
}
function sendAESKey(data) {
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            
        }
    };
    xhttp.open("POST", "/key/register/?id=", true);
    xhttp.send(data);
}

function pingSovy() {
    let xhttp = new XMLHttpRequest();
    let data = 
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
        
        }
    };
    xhttp.open("POST", "/key/ping/?id="+state.user.id, true);
    xhttp.send();
}

function importPublicKey(data) {
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            state.scrypt.serverPub = PublicFromXml(this.responseText);

        } else {
            //alert("failed to import key");
            reportState("failed to import RSA key");
        }
    };
    xhttp.open("GET", "/key/new/?id="+state.user.id, true);
    xhttp.send(); 
}
function encryptRSA(data) {
    return state.scrypt.serverPub.encript(data);
}

function encryptAES(params) {
    
}

function decryptAES(params) {
    
}

function prepareAES(params) {
    
}

function reportState(params) {
    document.getElementById("state").innerHTML = params;
}

function PublicToXml(data) {
    let x = pki.publicKeyToPem(key.publicKey);
    return "<key>"+x+"</key>";
}

function PublicFromXml(data) {
    let xmlDoc = parser.parseFromString(data,"text/xml");
    let y = xmlDoc.getElementsByTagName("key").valu;
    let x = pki.publicKeyFromPem(y);
    return x;
}

function getRandomString(length) {
    let out = "";
    let x = 0;
    for (let i = 0; i < length; i++) {
        out = out + charset[Math.random(charset.length)];
    }
    return out;
}

class State {
    constructor(){
        this.user = new User;
        this.scrypt = new SCrypt;
    }

    loadState(){
        if (typeof(Storage) !== "undefined") {
            let data = localStorage.getItem("sovygo_state");
            if (data == null) {
                
                return false;
            }
            if (data.outdated) {
                return false;
            } else {
                return data;
            }
        } else {
            return false;
        }
    }

    saveState(){
        if (typeof(Storage) !== "undefined") {
            let data = {
                outdated : false,
                user : {
                    username : this.user.name,
                    id : this.user.id
                },
                keys : {
                    //clientPub : pki.publicKeyToPem(this.scrypt.keyClient.publicKey),
                    //clientPri : pki.privateKeyToPem(this.scrypt.keyClient.privateKey),
                    serverPub : pki.publicKeyToPem(this.scrypt.keyServer),
                    symmetricKey : this.scrypt.symmetricKey,
                    symmetricIV : this.scrypt.symmetricIV
                }
            };
            localStorage.setItem();
        }
    }

    clearState(){
        if (typeof(Storage) !== "undefined") {
            localStorage.removeItem("sovygo_state");
        }
    }
}

class User {
    constructor(){
        this.username = "";
        this.id = getRandomString(16);
    }
}

class SCrypt {
    constructor(){
        //this.keyClient = forge.rsa.generateKeyPair();
        this.keyServer = null;
        this.symmetricKey = null;
        this.symmetricIV = null;
    }

    setServerKey(data) {
        this.keyServer = pki.publicKeyFromPem(data);
    }

}

function aplyState(data) {
    state.user.name = data.user.username;
    state.user.id = data.user.id;
    state.scrypt.setServerKey(data.keys.serverPub);
    state.scrypt.symmetricKey = data.keys.symmetricKey;
    state.scrypt.symmetricIV = data.keys.symmetricIV;

}

function newState(params) {
    state.user.id = getRandomString(16);
    importPublicKey();
}
setup();

setInterval(function(){
    pingSovy();
} ,250000);