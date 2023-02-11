$('#post-form').on('submit', createPost);
$('.like-post').on('click', likePost);
$('#update-post').on('click', updatePost);


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

async function likePost(e) {
    e.preventDefault();

    const clickedElement = $(e.target);
    const postId = clickedElement.closest('div').data('post-id');

    clickedElement.prop('disabled', true);
    clickedElement.addClass('disabled');

    $.ajax({
        url: '/posts/' + postId + '/like',
        method: 'POST',
        success: function(data) {
            const likesCounter = clickedElement.next('span')
            const likesCount = parseInt(likesCounter.text());
            likesCounter.text(likesCount + 1);
        },
        error: function(err) {
            console.log(err);
        },
        complete: function() {
            clickedElement.prop('disabled', false);
        }
    });
}

function updatePost(e) {
    $(this).prop('disabled', true);

    const postId = $(this).data('post-id');

    $.ajax({
        url: '/posts/' + postId,
        method: 'PUT',
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        },
        success: function(data) {
            window.location = '/home'
        },
        error: function(err) {
            console.log(err);
        },
        complete: function() {
            $('#update-post').prop('disabled', false);
        }
    });
}