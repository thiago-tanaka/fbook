$('#registration-form').on('submit', createUsers);

function createUsers(event) {
    event.preventDefault();
    const password = $('#password').val();

    if(password !== $('#confirm-password').val()) {
        Swal.fire({
            title: 'Error!',
            text: 'Passwords do not match!',
            icon: 'error',
        })
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
            Swal.fire({
                title: 'Success!',
                text: 'User created successfully!',
                icon: 'success',
            }).then(() => {
                $.ajax({
                    url: '/login',
                    type: 'POST',
                    data: {
                        email: $('#email').val(),
                        password: $('#password').val()
                    },
                    success: function (data) {
                        window.location.href = '/home';
                    },
                    error: function (err) {
                        Swal.fire({
                            title: 'Error!',
                            text: 'User could not be created!',
                            icon: 'error',
                        })
                    }
                })

            })
        },
        error: function (err) {
            Swal.fire({
                title: 'Error!',
                text: 'User could not be created!',
                icon: 'error',
            })
        }
    });
}