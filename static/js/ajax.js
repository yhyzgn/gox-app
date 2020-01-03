(function () {
    if (typeof window.$ !== "undefined") {
        $.plugin("ajax", function (arg) {
            let xhr = new XMLHttpRequest();
            xhr.open(arg.method, arg.url, true);
            if (arg.headers !== undefined) {
                for (let name in arg.headers) {
                    xhr.setRequestHeader(name, arg.headers[name])
                }
            }
            xhr.send();

            xhr.onreadystatechange = function () {
                if (xhr.readyState === 4) {
                    if (xhr.status === 200 && arg.success !== undefined) {
                        arg.success(xhr.responseText);
                    } else if (arg.error !== undefined) {
                        arg.error(xhr.response);
                    }
                }
            };

            xhr.onerror = function (e) {
                console.log(e)
            }
        })
    }
})();