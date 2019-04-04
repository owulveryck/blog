var canvas;
var image;
var bThick = 20;
var cColour = '#FFFFFF'; // #black

//Initialize  fabric 
function initCanvas(){
    canvas = new fabric.Canvas('canvasBox');
    canvas.isDrawingMode = true;
    canvas.freeDrawingBrush.width = bThick;
    canvas.freeDrawingBrush.color = cColour;
    canvas.backgroundColor = '#000000';            
}

window.onload = function(){
    initCanvas();
    // submit button
    var btnsub = document.getElementById('btnSubmit');
    btnsub.addEventListener('click', postImg());
};

//reset canvas
function reset() {           
    canvas.clear();   
    $('#guess').text('');     
}

function postImg(){                        
    var imgURL = canvas.toDataURL();   

    //Send Ajax call
    $.ajax({
        type: 'post',
        url: '/picture',
        data: imgURL,
    
        success: function(data){
            $('#guess').text(data);
        }              
    });                     
}

if (!WebAssembly.instantiateStreaming) { // polyfill
         WebAssembly.instantiateStreaming = async (resp, importObject) => {
                 const source = await (await resp).arrayBuffer();
                 return await WebAssembly.instantiate(source, importObject);
       };
}

  const go = new Go();
  let mod, inst;
async function load() {
  document.getElementById("loadButton").disabled = true;
  WebAssembly.instantiateStreaming(fetch("/wasm/nnre.wasm"), go.importObject).then((result) => {
         mod = result.module;
         inst = result.instance;
         document.getElementById("runButton").disabled = false;
  });
}

async function run() {
       console.clear();
       await go.run(inst);
       inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
}
function loadFile(){
    var x = document.getElementById("knowledgeFile");
    var txt = "";
    if ('files' in x) {
        if (x.files.length == 0) {
            txt = "Select one or more files.";
        } else {
            for (var i = 0; i < x.files.length; i++) {
                txt += "<br><strong>" + (i+1) + ". file</strong><br>";
                var file = x.files[i];
                if ('name' in file) {
                    txt += "name: " + file.name + "<br>";
                }
                if ('size' in file) {
                    txt += "size: " + file.size + " bytes <br>";
                }
            }
        }
    }
    else {
        if (x.value == "") {
            txt += "Select one or more files.";
        } else {
            txt += "The files property is not supported by your browser!";
            txt  += "<br>The path of the selected file: " + x.value; // If the browser does not support the files property, it will return the path of the selected file instead.
        }
    }
    document.getElementById("demo").innerHTML = txt;
}
