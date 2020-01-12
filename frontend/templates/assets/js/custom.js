function callDetail(id) {
    let address = "/detail?id=" + id;
    window.location = encodeURI(address);
}

function deleteActivity(id) {
    let answer = confirm("Soll die Aktivität wirklich gelöscht werden?")
    if (answer) {
        let address = "/delete?id=" + id;
        window.location = encodeURI(address)
    }
}