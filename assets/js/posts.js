$('#post-form').on('submit', createPost);


function createPost(e) {
    e.preventDefault();

    $.ajax({
        url: '/posts',
        method: 'POST',
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        },
        success: function(data) {
            window.location = '/home'
        },
        error: function(err) {
            console.log(err);
        }
    });
}