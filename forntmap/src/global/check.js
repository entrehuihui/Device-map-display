
export default {
    checkPhone,
    checkEmail,
    checkName,
    checkPassword
}

function checkPhone(phone) {
    if ((/^1[34578]\d{9}$/.test(phone))) {
        return true;
    }
    return false;
}

function checkEmail(email) {
    var reg = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
    if (reg.test(email)) {
        return true;
    }
    return false;
}

function checkName(name) {
    //用户名正则，字母开头,4到16位（字母，数字，下划线，减号）
    var uPattern = /^[a-zA-Z]{1}[a-zA-Z0-9_-]{3,16}$/;
    if (uPattern.test(name)) {
        return true
    }
    return false
}

function checkPassword(password) {
    if (password.length < 6) {
        return false
    }
    return true
}