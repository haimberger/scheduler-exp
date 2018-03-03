$(document).ready(function() {
  $('.task-title').click(function() {
    $(this)
      .parents('.task')
      .find('.task-menu')
      .toggle();
  });
});
