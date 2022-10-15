function PostConf(conf) {
    var http = new XMLHttpRequest();
    var url = 'http://localhost:8888/conf';
    http.open('POST', url, true);

    //Send the proper header information along with the request
    http.setRequestHeader('Content-type', 'application/json');

    http.onreadystatechange = function () {//Call a function when the state changes.
        if (http.readyState == 4 && http.status == 200) {
            alert(http.responseText);
        }
    }
    http.send(conf);
}

function GetConf() {
    console.log("Get Conf")
    var http = new XMLHttpRequest();
    var url = 'http://localhost:8888/conf';
    http.open('GET', url, true);

    http.onreadystatechange = function () {//Call a function when the state changes.
        if (http.readyState == 4 && http.status == 200) {
            Get(JSON.parse(http.responseText))
        }
    }
    http.send();
}

function BuildArea(Path, Format, Default, num) {
    let area = document.createElement('div')
    area.className = "area"

    let name = document.createElement('p')
    name.className = "formatTitle"
    name.innerHTML = "Format " + num

    let col1 = document.createElement('div')
    col1.innerHTML = "<label>Chemin vers le texte</label><label>Format du texte</label><label>Texte si pas de musique</label>"
    col1.className = "col"
    let col2 = document.createElement('div')
    col2.className = "col"

    col2.appendChild(BuildInput("output/test.txt", "inputFieldTxtPath", Path))
    col2.appendChild(BuildInput("%artist% - %title% (%year%)", "inputFieldFormat", Format))
    col2.appendChild(BuildInput("No song is currently playing", "inputFieldDefault", Default))

    area.appendChild(name)
    area.appendChild(col1)
    area.appendChild(col2)

    return area
}

function BuildInput(placeholder, name, value) {
    var input = document.createElement('input')
    input.type = "text"
    input.className = "inputField field"
    input.name = name
    input.placeholder = placeholder
    input.value = value
    return input
}

function Get(fileStruct) {
    document.getElementsByName("inputFieldFrequency")[0].value = fileStruct.Frequency / 1000000000
    document.getElementsByName("inputFieldIsImgSaved")[0].checked = fileStruct.SaveImg
    document.getElementsByName("inputFieldImgPath")[0].value = fileStruct.ImgPath

    let formats = document.getElementById("formats")
    while (formats.childElementCount > 0) {
        formats.removeChild(formats.lastChild)
    }

    for (let i = 0; i < fileStruct.FileSaves.length; i++) {
        let area = BuildArea(fileStruct.FileSaves[i].Path, fileStruct.FileSaves[i].Format, fileStruct.FileSaves[i].Default, i + 1)
        formats.appendChild(area)
    }
}

function Save() {
    let fileStruct = {
        Frequency: "",
        SaveImg: "",
        ImgPath: "",
        FileSaves: []
    }
    fileStruct.Frequency = document.getElementsByName("inputFieldFrequency")[0].value
    if (fileStruct.Frequency == "") {
        alert("Frequency is empty")
        return
    }
    fileStruct.Frequency = Math.round(parseFloat(fileStruct.Frequency) * 1000000000)
    fileStruct.SaveImg = document.getElementsByName("inputFieldIsImgSaved")[0].checked
    fileStruct.ImgPath = document.getElementsByName("inputFieldImgPath")[0].value
    if (fileStruct.ImgPath == "") {
        alert("Image path is empty")
        return
    }

    let areas = document.getElementById("formats").getElementsByClassName("area")

    for (let i = 0; i < areas.length; i++) {
        let col = areas[i].childNodes[2]
        let elem = {
            Path: col.childNodes[0].value,
            Format: col.childNodes[1].value,
            Default: col.childNodes[2].value
        }
        console.log(elem)
        if (elem.Path == "") {
            alert("One of the text path is empty")
            return
        }
        fileStruct.FileSaves.push(elem)
    }

    PostConf(JSON.stringify(fileStruct))
}

function AddFormat() {
    let formats = document.getElementById("formats")
    var area = BuildArea("", "", "", formats.childElementCount + 1)
    formats.appendChild(area)
}

function RemoveFormat() {
    let formats = document.getElementById("formats")
    if (formats.childElementCount > 0) {
        formats.removeChild(formats.lastChild)
    }
}

/* =========
 *   Theme 
 * =========*/

let theme = [
    {
        bodyBackground: "linear-gradient(160deg, rgba(20, 13, 43, 1) 0%, rgba(30, 17, 70, 1) 15%, rgba(33, 15, 80, 1) 35%, rgba(37, 10, 92, 0.9651972157772621) 59%, rgba(54, 9, 121, 0.9466357308584686) 71%, rgba(190, 58, 146, 0.9350348027842228) 83%, rgba(0, 239, 255, 1) 100%)",
        titleColor: "#d845a7",
        categoryTitleColor: "#e287c4",
        buttonChangeBackground: "white",
        buttonSaveBackground: "white",
        fieldBackground: "white",
        checkColor: "#ce78b1",
        boxShadow: "#d845a731",
        box1: "",
        box2: "",
        box3: "",
        box4: "",
        box5: ""
    },
    {
        bodyBackground: "linear-gradient(90deg, rgba(131,58,180,1) 0%, rgba(253,29,29,1) 50%, rgba(252,176,69,1) 100%)",
        titleColor: "yellow",
        categoryTitleColor: "#ffbe00",
        buttonChangeBackground: "linear-gradient(135deg, rgba(255,145,0,1) 0%, rgba(255,209,0,1) 30%, rgba(255,252,0,1) 100%)",
        buttonSaveBackground: "linear-gradient(135deg, rgba(255,145,0,1) 0%, rgba(255,209,0,1) 30%, rgba(255,252,0,1) 100%)",
        fieldBackground: "",
        checkColor: "#e91b16",
        boxShadow: "",
        box1: "",
        box2: "",
        box3: "",
        box4: "",
        box5: ""
    },
    {
        bodyBackground: "linear-gradient(180deg, rgba(99,159,171,1) 43%, rgba(34,34,34,1) 100%)",
        titleColor: "white",
        categoryTitleColor: "#6a6969",
        buttonChangeBackground: "white",
        buttonSaveBackground: "white",
        fieldBackground: "#b5b5b5",
        checkColor: "#6a6969",
        boxShadow: "",
        box1: "",
        box2: "",
        box3: "",
        box4: "",
        box5: ""
    },
    {
        bodyBackground: "linear-gradient(90deg, rgba(91,192,235,1) 0%, rgba(253,231,76,1) 25%, rgba(155,197,61,1) 50%, rgba(195,66,63,1) 75%, rgba(33,26,30,1) 100%)",
        titleColor: "red",
        categoryTitleColor: "#211A1E",
        buttonChangeBackground: "",
        buttonSaveBackground: "",
        fieldBackground: "",
        checkColor: "#211A1E",
        boxShadow: "#cdcdcd",
        box1: "#5BC0EB",
        box2: "#FDE74C",
        box3: "#9BC53D",
        box4: "#C3423F",
        box5: "#211A1E",
    },
    {
        bodyBackground: "linear-gradient(90deg, rgba(91,192,235,1) 0%, rgba(253,231,76,1) 25%, rgba(155,197,61,1) 50%, rgba(195,66,63,1) 75%, rgba(33,26,30,1) 100%)",
        titleColor: "red",
        categoryTitleColor: "#211A1E",
        buttonChangeBackground: "",
        buttonSaveBackground: "",
        fieldBackground: "",
        checkColor: "#211A1E",
        boxShadow: "",
        box1: "",
        box2: "",
        box3: "",
        box4: "",
        box5: "",
    },
    {
        bodyBackground: "linear-gradient(163deg, rgba(250,255,3,1) 0%, rgba(29,253,239,1) 24%, rgba(232,69,252,1) 85%)",
        titleColor: "",
        categoryTitleColor: "",
        buttonChangeBackground: "",
        buttonSaveBackground: "",
        fieldBackground: "",
        checkColor: "",
        boxShadow: "",
        box1: "",
        box2: "",
        box3: "",
        box4: "",
        box5: ""
    }
]

let curTheme = 0

function ChangeColor() {
    curTheme = (curTheme + 1) % theme.length
    PostTheme()
    t = theme[curTheme]

    document.querySelector(":root").style.setProperty("--body-background", t.bodyBackground)
    document.querySelector(":root").style.setProperty("--title-color", t.titleColor)
    document.querySelector(":root").style.setProperty("--category-title-color", t.categoryTitleColor)
    document.querySelector(":root").style.setProperty("--button-change-background", t.buttonChangeBackground)
    document.querySelector(":root").style.setProperty("--button-save-background", t.buttonSaveBackground)
    document.querySelector(":root").style.setProperty("--field-background", t.fieldBackground)
    document.querySelector(":root").style.setProperty("--check-color", t.checkColor)
    document.querySelector(":root").style.setProperty("--box-shadow", t.boxShadow)
    document.querySelector(":root").style.setProperty("--box1", t.box1)
    document.querySelector(":root").style.setProperty("--box2", t.box2)
    document.querySelector(":root").style.setProperty("--box3", t.box3)
    document.querySelector(":root").style.setProperty("--box4", t.box4)
    document.querySelector(":root").style.setProperty("--box5", t.box5)
}

function PostTheme() {
    var http = new XMLHttpRequest();
    var url = 'http://localhost:8888/theme?num=' + curTheme.toString();
    http.open('POST', url, true);

    http.onreadystatechange = function () {//Call a function when the state changes.
        if (http.readyState == 4 && http.status == 200) {
            console.log(http.responseText);
        }
    }
    http.send();
}

function GetTheme() {
    console.log("Get Theme")
    var http = new XMLHttpRequest();
    var url = 'http://localhost:8888/theme';
    http.open('GET', url, true);

    http.onreadystatechange = function () {//Call a function when the state changes.
        if (http.readyState == 4 && http.status == 200) {
            curTheme = parseInt(http.responseText)
            curTheme--
            ChangeColor()
        }
    }
    http.send();
}

function GetAll() {
    GetConf()
    GetTheme()
}

document.onload = GetAll()