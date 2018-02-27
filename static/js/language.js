$(document)
    .ready(function() {
        function setCookie(cname, cvalue, exdays) {
            var d = new Date();
            d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
            var expires = "expires="+d.toUTCString();
            document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
        }


        $('#select_language').dropdown({
            action: function (text, value) {
                setCookie('lang', value, 365);
                window.location.reload();
            }
        });
    });