const success_text = '' +
    '<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="bi bi-binoculars-fill" viewBox="0 0 16 16">\n' +
    '  <path d="M4.5 1A1.5 1.5 0 0 0 3 2.5V3h4v-.5A1.5 1.5 0 0 0 5.5 1h-1zM7 4v1h2V4h4v.882a.5.5 0 0 0 .276.447l.895.447A1.5 1.5 0 0 1 15 7.118V13H9v-1.5a.5.5 0 0 1 .146-.354l.854-.853V9.5a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5v.793l.854.853A.5.5 0 0 1 7 11.5V13H1V7.118a1.5 1.5 0 0 1 .83-1.342l.894-.447A.5.5 0 0 0 3 4.882V4h4zM1 14v.5A1.5 1.5 0 0 0 2.5 16h3A1.5 1.5 0 0 0 7 14.5V14H1zm8 0v.5a1.5 1.5 0 0 0 1.5 1.5h3a1.5 1.5 0 0 0 1.5-1.5V14H9zm4-11H9v-.5A1.5 1.5 0 0 1 10.5 1h1A1.5 1.5 0 0 1 13 2.5V3z"/>\n' +
    '</svg>\n' +
    'Start a test';

const danger_test = '' +
    '<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="bi bi-stopwatch" viewBox="0 0 16 16">\n' +
    '  <path d="M8.5 5.6a.5.5 0 1 0-1 0v2.9h-3a.5.5 0 0 0 0 1H8a.5.5 0 0 0 .5-.5V5.6z"/>\n' +
    '  <path d="M6.5 1A.5.5 0 0 1 7 .5h2a.5.5 0 0 1 0 1v.57c1.36.196 2.594.78 3.584 1.64a.715.715 0 0 1 .012-.013l.354-.354-.354-.353a.5.5 0 0 1 .707-.708l1.414 1.415a.5.5 0 1 1-.707.707l-.353-.354-.354.354a.512.512 0 0 1-.013.012A7 7 0 1 1 7 2.071V1.5a.5.5 0 0 1-.5-.5zM8 3a6 6 0 1 0 .001 12A6 6 0 0 0 8 3z"/>\n' +
    '</svg>' +
    'Wait';

const URL = '127.0.0.1:5000/test';
let flag = true;

function request() {
    fetch(URL)
        .then(async (response) => {
            document.getElementById("result").innerText = await response.json();
        })
        .catch((error) => {
            document.getElementById("result").innerText = "Problem in testing ..."
            console.error(error)
        })

    setTimeout(reset, 2000)
}


function begin() {
    if (!flag)
        return;

    flag = false;

    const el = document.getElementById("main-btn")

    el.classList.remove("btn-success")
    el.classList.add("btn-danger")
    el.innerHTML = danger_test;

    request()
}

function reset() {
    flag = true;

    const el = document.getElementById("main-btn")

    el.classList.remove("btn-danger")
    el.classList.add("btn-success")
    el.innerHTML = success_text;
}