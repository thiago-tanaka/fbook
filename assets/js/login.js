$('#login').on('submit', function (e) {
    e.preventDefault();
    var user = {
        email: $('#email').val(),
        password: $('#password').val()
    };
    $.ajax({
        url: '/login',
        type: 'POST',
        data: user,
        success: function (data) {
            window.location.href = '/home';
        },
        error: function (err) {
            console.log(err);
        }
    });
});