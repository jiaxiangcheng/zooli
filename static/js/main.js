


$(document)
    .ready(function() {

        var sideBarIsHide = false;
        var ManuelSideBarIsHide = false;
        var ManuelSideBarIsState = false;


        $(window).resize(function() {
            if (ManuelSideBarIsHide == false) {
                if ($(window).width() <= 767) {
                    if (!sideBarIsHide); {
                        resizeSidebar("1");
                        sideBarIsHide = true;
                        $(".colhidden").addClass("displaynone");

                    }
                } else {
                    if (sideBarIsHide); {
                        resizeSidebar("0");
                        sideBarIsHide = false;

                        $(".colhidden").removeClass("displaynone");

                    }
                }
            }
        });


        function resizeSidebar(op) {
            if (op == "1") {

                $(".ui.sidebar.left").addClass("very thin icon");
                $(".navslide").addClass("marginlefting");
                $(".sidebar.left span").addClass("displaynone");
                $(".sidebar .accordion").addClass("displaynone");
                $(".ui.dropdown.item.displaynone").addClass("displayblock");
                $($(".logo img")[0]).addClass("displaynone");
                $($(".logo img")[1]).removeClass("displaynone");
                $(".hiddenCollapse").addClass("displaynone");


            } else {

                $(".ui.sidebar.left").removeClass("very thin icon");
                $(".navslide").removeClass("marginlefting");
                $(".sidebar.left span").removeClass("displaynone");
                $(".sidebar .accordion").removeClass("displaynone");
                $(".ui.dropdown.item.displaynone").removeClass("displayblock");
                $($(".logo img")[1]).addClass("displaynone");
                $($(".logo img")[0]).removeClass("displaynone");
                $(".hiddenCollapse").removeClass("displaynone");


            }

        }

        $("#full_screen_toggle").on('click', function () {
            if (!document.fullscreenElement &&    // alternative standard method
                !document.mozFullScreenElement && !document.webkitFullscreenElement && !document.msFullscreenElement ) {  // current working methods
                if (document.documentElement.requestFullscreen) {
                    document.documentElement.requestFullscreen();
                } else if (document.documentElement.msRequestFullscreen) {
                    document.documentElement.msRequestFullscreen();
                } else if (document.documentElement.mozRequestFullScreen) {
                    document.documentElement.mozRequestFullScreen();
                } else if (document.documentElement.webkitRequestFullscreen) {
                    document.documentElement.webkitRequestFullscreen(Element.ALLOW_KEYBOARD_INPUT);
                }
            } else {
                if (document.exitFullscreen) {
                    document.exitFullscreen();
                } else if (document.msExitFullscreen) {
                    document.msExitFullscreen();
                } else if (document.mozCancelFullScreen) {
                    document.mozCancelFullScreen();
                } else if (document.webkitExitFullscreen) {
                    document.webkitExitFullscreen();
                }
            }
        });


        $("#nav_sidebar .item").on('click', function (event) {
            event.preventDefault();
            $.get(this.href, {}, function (response) {
                $('#main_content').html(response);
            });
        });

        $("#open_btn").on("click", function() {

            ManuelSideBarIsHide = true;
            if (!ManuelSideBarIsState) {
                resizeSidebar("1");
                ManuelSideBarIsState = true;
            } else {
                resizeSidebar("0");
                ManuelSideBarIsState = false;
            }
        });

    });