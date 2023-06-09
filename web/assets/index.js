const API_URL = "http://localhost:8080";

const updateUrlID = (id) => {
    let url = new URL(window.location);
    url.searchParams.set("id", id);
    window.history.pushState({}, "", url);
}

const printError = (error) => {
    document.querySelector("#response").style.background = "#ff1818a1";
    document.querySelector("#responseText").innerHTML = error;
}

const printResponse = (responseText, responseId, responseToken, ResponseTime) => {
    document.querySelector("#response").style.background = "#fff518a1";
    updateUrlID(responseId);
    showReportBtn();
    document.querySelector("#responseText").innerHTML = responseText;
    document.querySelector("#details #responseId").value = responseId;
    document.querySelector("#details #responseCredit").value = responseToken;
    document.querySelector("#details #responseTime").value = ResponseTime;
}

const toggleButtons = (toggle) => {
    document.querySelectorAll('button').forEach(e => {
        e.disabled = toggle;
    });
}

const showReportBtn = () => {
    document.querySelector("#reportBtn").disabled = false;
}

const copyFunction = () => {
    let responseText = document.querySelector("#responseText");
    navigator.clipboard.writeText(responseText.innerText);

    document.querySelector("#copyBtn").innerHTML = "copied";
    setTimeout(() => {
        document.querySelector("#copyBtn").innerHTML = "copy";
    }, 500);
}

const reportFunction = () => {
    document.querySelector("#reportBtn").innerHTML = "reported";
    document.querySelector("#reportBtn").disabled = true;

    const responseId = document.querySelector("#details #responseId").value;

    let request = APIRequest("/report", "POST", {
        id: responseId
    });
    request.then((data) => {
        if (data.ErrorCode) {
            printError(data.ErrorMessage);
            return;
        }

        document.querySelector("#reportBtn").innerHTML = data.ResponseText;
    });
}

const Loading = () => {
    document.querySelector("#reportBtn").innerHTML = "report";
    document.querySelector("#responseText").innerHTML = "loading...";
    toggleButtons(true);
}
const LoadingDone = () => {
    document.querySelector("#responseText").innerHTML = "";
    toggleButtons(false);
}

const APIRequest = (url, method, body) => {
    return fetch(API_URL + url, {
        method: method,
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
    })
        .then(response => response.json())
        .catch(error => printError(error));
}

const LoadLatests = () => {
    let request = APIRequest("/latest", "GET");
    request.then((data) => {
        if (data.ErrorCode) {
            printError(data.ErrorMessage);
            return;
        }

        document.querySelector("#latest").innerHTML = data.ResponseText;
        document.querySelectorAll('#latest li span').forEach(e => {
            let d = new Date(e.innerText * 1000);
            e.innerText = d.getDate() + '/' + (d.getMonth() + 1) + '/' + d.getFullYear().toString().slice(-2) + " " + String(d.getHours()).padStart(2, '0') + ':' + String(d.getMinutes()).padStart(2, '0');
        });

        document.querySelector("#refreshBtn").disabled = false;
    });
};

const RefreshLatests = () => {
    document.querySelector("#refreshBtn").disabled = true;
    LoadLatests();
};

document.querySelector("form").addEventListener("submit", (event) => {
    event.preventDefault();
    Loading();

    const message = document.querySelector("#message").value;

    let request = APIRequest("/generate", "POST", {
        desc: message
    });
    request.then((data) => {
        LoadingDone();

        if (data.ErrorCode) {
            printError(data.ErrorMessage);
            return;
        }

        printResponse(data.ResponseText, data.ResponseID, data.ResponseToken, data.ResponseTime);
    });
});
document.querySelector("#copyBtn").addEventListener("click", () => {
    copyFunction();
});
document.querySelector("#reportBtn").addEventListener("click", () => {
    reportFunction();
});
document.querySelector("#refreshBtn").addEventListener("click", () => {
    RefreshLatests();
});
document.querySelector("#showBtn").addEventListener("click", () => {
    const value = document.querySelector('#responseText').innerText;
    var encoded_value = encodeURIComponent(value);

    if (encoded_value.startsWith("/")) encoded_value = encoded_value.substring(1);
    if (encoded_value.endsWith("/")) encoded_value = encoded_value.substring(0, encoded_value.length - 1);

    window.open('https://regex101.com/?regex=' + encoded_value, '_blank');
});

let queryParams = new URLSearchParams(window.location.search);
if (queryParams.get("id") != null) {
    Loading();

    let request = APIRequest("/get", "POST", {
        id: queryParams.get("id")
    });
    request.then((data) => {
        LoadingDone();

        if (data.ErrorCode) {
            printError(data.ErrorMessage);
            return;
        }

        document.querySelector("#message").value = data.InputText;
        printResponse(data.ResponseText, data.ResponseID, data.ResponseToken, data.ResponseTime);
    });
}

RefreshLatests();