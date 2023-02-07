function getCookie(name) {
    var cookieArr = document.cookie.split(";");
    console.log(cookieArr);
    for(var i = 0; i < cookieArr.length; i++) {
        var cookiePair = cookieArr[i].split("=");
        if(name == cookiePair[0].trim()) {
            return decodeURIComponent(cookiePair[1]);
        }
    }
    return null;
}

function checkAuthCookie() {
    var token = getCookie("access_token");
    cobsole.log(token);

    if (!token) {
        window.location.href = "/login";
    }
}

checkAuthCookie();