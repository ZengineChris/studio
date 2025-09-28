import "./../css/main.css"
import "htmx.org"
import "alpinejs"

function onScanSuccess(decodedText, decodedResult) {
    // handle the scanned code as you like, for example:
    console.log(`Code matched = ${decodedText}`, decodedResult);
}

function onScanFailure(error) {
    console.warn(`Code scan error = ${error}`);
}

let html5QrcodeScanner = new Html5QrcodeScanner(
    "reader",
    { fps: 10, qrbox: { width: 250, height: 250 } },
  /* verbose= */ false);
html5QrcodeScanner.render(onScanSuccess, onScanFailure);

function test() {
    Html5Qrcode.getCameras().then(devices => {
        console.log("hello")
        console.log(devices)
        if (devices && devices.length) {
            console.log(devices)
            var cameraId = devices[0].id;
        }
    }).catch(err => {
        console.error(err)
    });
}

