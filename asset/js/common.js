function toTop() {
    $(window).scroll(function () {
        if ($(window).scrollTop() > 200) {
            $('.to-top').fadeIn(500);
        } else {
            $('.to-top').fadeOut(500)
        }
    });

    $("#back-to-top").click(function () {
        $('body,html').animate({scrollTop: 0}, 200);
        return false;
    });
}