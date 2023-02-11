$(document).ready(function() {
    $('#post-form').on('submit', createPost);
    $(document).on('click', '.like-post', toggleLikeDislike);
    $(document).on('click', '.dislike-post', toggleLikeDislike);
    $('#update-post').on('click', updatePost);
    $('.delete-post').on('click', deletePost);
});

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
            Swal.fire({
                title: 'Error!',
                text: 'Post could not be created',
                icon: 'error',
            })
        }
    });
}

async function toggleLikeDislike(e) {
    e.preventDefault();

    const $clickedElement = $(e.target);
    const postId = $clickedElement.closest('div').data('post-id');
    const isLike = $clickedElement.hasClass('like-post');
    const url = isLike ? '/posts/' + postId + '/like' : '/posts/' + postId + '/dislike';
    const method = 'POST';

    $clickedElement.prop('disabled', true);

    $.ajax({
        url: url,
        method: method,
        success: function(data) {
            const $likesCounter = $clickedElement.next('span');
            let likesCount = parseInt($likesCounter.text());

            if (isLike) {
                likesCount++;
                $clickedElement
                    .removeClass('like-post')
                    .addClass('dislike-post')
                    .addClass('text-danger');
            } else {
                likesCount--;
                $clickedElement
                    .removeClass('dislike-post')
                    .removeClass('text-danger')
                    .addClass('like-post');
            }

            $likesCounter.text(likesCount);
        },
        error: function(err) {
            Swal.fire({
                title: 'Error!',
                text: 'Post could not be updated',
                icon: 'error',
            })
        },
        complete: function() {
            $clickedElement.prop('disabled', false);
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
            Swal.fire({
                title: 'Successo!',
                text: "Post Updated Successfully!",
                icon: 'success',
            }).then(() => {
                window.location = '/home'
            })
            // window.location = '/home'
        },
        error: function(err) {
            Swal.fire({
                title: 'Error!',
                text: 'Post could not be updated',
                icon: 'error',
            })
        },
        complete: function() {
            $('#update-post').prop('disabled', false);
        }
    });
}

function deletePost(e) {
    $(this).prop('disabled', true);

    Swal.fire({
        title: 'Are you sure?',
        text: "You won't be able to revert this!",
        icon: 'warning',
        showCancelButton: true,
    }).then((confirmation) => {
        if (!confirmation.value) {
            return
        }

        const postId = $(this).data('post-id');

        $.ajax({
            url: '/posts/' + postId,
            method: 'DELETE',
            success: function(data) {
                window.location = '/home'
            },
            error: function(err) {
                Swal.fire({
                    title: 'Error!',
                    text: 'Post could not be deleted',
                    icon: 'error',
                })
            },
            complete: function() {
                $('#delete-post').prop('disabled', false);
            }
        });
    })


}