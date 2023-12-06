async function encrypt() {
    let f_enc_type = document.getElementById("enc-type");
    let enc_type = f_enc_type.value;

    let f_pub_key = document.getElementById("pub-key");
    let pub_key = f_pub_key.value;

    let f_enc_data = document.getElementById("enc-data");
    let enc_data = f_enc_data.value;

    if (enc_type.length === 0)
        alert("Encryption Type is required to encrypt.");
    else if (pub_key.length === 0)
        alert("Public Key is required to encrypt.");
    else if (enc_data.length === 0)
        alert("Encryption Data is required to encrypt.");

    let responseText = await getRequest(
        url = "http://localhost:8082/encrypt",
        method = "POST",
        body = {
            encryptionType: enc_type,
            publicKey: pub_key,
            data: enc_data,
        }
    );

    let f_enc_result = document.getElementById("enc-result");
    f_enc_result.value = responseText;
}

async function decrypt() {

}

async function getRequest(url = "", method = "GET", body = undefined) {
    let response = await fetch(url, {
        method: method,
        mode: "cors",
        body: JSON.stringify(body),
    });

    return response.text();
}
