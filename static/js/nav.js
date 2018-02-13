$(document)
    .ready(function() {
        $("#nav_menu .item").bind('click', function (event) {
            event.preventDefault();
            $.get(this.href, {}, function (response) {
                $('#content').html(response);
            });
        });
    });