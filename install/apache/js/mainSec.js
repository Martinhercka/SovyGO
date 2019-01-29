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
    let toSend = new Envelop;
    //toSend.body = state.scrypt.symmetricKey;
    toSend.RSAToEnvelop(state.scrypt.symmetricKey)
    let req = toSend.buildEnvelop();
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var env = new Envelop;
            env.fromEnvelop(this.responseText);
            console.log(env.body);
            
        }
    };
    xhttp.open("POST", "http://itsovy.sk:1122/key/aes/", true);
    console.log(req);
    xhttp.send(req);
}

function pingSovy() {
    let xhttp = new XMLHttpRequest();
    let data = new Envelop;
    data.encryptToEnvelop("ping")
    let req = data.buildEnvelop();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
        
        } else if (this.readyState == 4 && this.status == 403) {
            state.clearState();
            state = null;
            setup();
        }
    };
    xhttp.open("POST", "http://itsovy.sk:1122/key/ping/", true);
    console.log(req);
    xhttp.send(req);
}

function importPublicKey(data) {
    let toSend = new Envelop;
    toSend.encryption = false;
    toSend.body = JSON.stringify({sessionid : state.user.id});
    let req = toSend.buildEnvelop();
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            console.log(this.responseText);
            let x = String(this.responseText);
            x = x.split("&#xA;").join("");
            x = x.split("\n").join("");
            console.log(x);
            state.scrypt.serverPub = pki.publicKeyFromPem(x);
            reportState("imported");
            sendAESKey();
        } else {
            //alert("failed to import key");
            reportState("failed to import RSA key");
        }
    };
    xhttp.open("POST", "http://itsovy.sk:1122/key/new/", true);
    console.log(req);
    xhttp.send(req); 
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
        out = out + charset[Math.floor(Math.random()*charset.length)];
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
                    serverPub : pki.publicKeyToPem(this.scrypt.serverPub),
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
        this.serverPub = null;
        this.symmetricKey = null;
        this.symmetricIV = null;
    }

    setServerKey(data) {
        this.serverPub = pki.publicKeyFromPem(data);
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
    state.scrypt.symmetricKey = forge.random.getBytesSync(32);
    importPublicKey();
}

setInterval(function(){
    pingSovy();
} ,250000);

class Envelop{
    constructor(){
        this.encryption = true;
        this.sessionid = state.user.id;
        this.body = {};
        this.key = "";
        this.check = "";
    }

    toEnvelop(params) {
        this.encryption = false;
    
    }

    encryptToEnvelop(params) {
        this.encryption = true;
        this.key = forge.random.getBytesSync(16);
        let cipher = forge.cipher.createCipher('AES-CBC', state.scrypt.symmetricKey);
        cipher.start({iv: this.key});
        cipher.update(forge.util.createBuffer(params));
        cipher.finish();
        this.body = cipher.output;
    }

    RSAToEnvelop(params) {
        this.encryption = true;
        if (params == null) {
            this.body = state.scrypt.serverPub.encrypt(this.body);
        } else {
            this.body = state.scrypt.serverPub.encrypt(params);
        }
    }

    fromEnvelop(params) {
        let xmlDoc = parser.parseFromString(params, "text/xml");
        this.encryption = xmlDoc.getElementsByTagName("encryption");
        this.body = xmlDoc.getElementsByTagName("Body");
        this.key = xmlDoc.getElementsByTagName("Key");
        this.check = xmlDoc.getElementsByTagName("Check");
        if (this.encryption == true) {
            let decipher = forge.cipher.createDecipher('AES-CBC', state.scrypt.symmetricKey);
            decipher.start({iv: this.key});
            decipher.update(this.body);
            this.body = decipher.finish();
        }
        return this.body;
    }

    buildEnvelop(){
        let out = "<Data>"+
            "\n\t<Head>"+
            "\n\t\t<encryption>"+this.encryption+"</encryption>"+
            "\n\t\t<session>"+this.sessionid+"</session>"+
            "\n\t</Head>"+
            "\n\t<Body>"+this.body+"</Body>"+
            "\n\t<Key>"+this.key+"</Key>"+
            "\n\t<Check>"+this.check+"</Check>"+
        "\n</Data>";
        return out;
    }
}


setup(); //leave on the end of file. if not it will cause lexical error