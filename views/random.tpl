<a class="item" href="/dont_use-long_name-and_bar" id="random_data">
    <i class="edit outline icon"></i> Random Data
</a>


<div class="ui basic modal" id="data_modal">
    <div class="ui icon header">
        <i class="archive icon"></i>
        Generating Random data set
    </div>
    <div class="content">
        <div class="ui segment">
            <div class="ui active dimmer">
                <div class="ui indeterminate huge text loader">Generating...</div>
            </div>
            <p>It may takes minutes</p>
        </div>
    </div>
</div>



<script>
    $(document)
            .ready(function() {
                $("#random_data").bind('click', function (event) {
                    event.preventDefault();
                    $('#data_modal').modal('show');
                    $.get(this.href, {}, function (response) {
                        $('#main_content').html(response);
                        $('#data_modal').modal('hide');
                    });
                });
            });

</script>