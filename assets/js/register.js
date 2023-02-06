$('#registration-form').on('submit', createUsers);

function createUsers(event) {
    event.preventDefault();
    const password = $('#password').val();

    if(password !== $('#confirm-password').val()) {
        alert('Passwords do not match');
        return;
    }

    var user = {
        name: $('#name').val(),
        email: $('#email').val(),
        nick: $('#nick').val(),
        password: password
    };

    $.ajax({
        url: '/users',
        type: 'POST',
        data: user,
        success: function (data) {
            console.log('success');
        },
        error: function (err) {
            console.log(err);
        }
    });
}