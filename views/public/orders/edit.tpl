<form class="ui form" enctype="multipart/form-data">
    <h2 id="title">
        <i class="payment icon"></i>
        Order Information
    </h2>
    {{template "public/orders/form/body.tpl" .}}
    <button id="save" class="ui primary button" type="submit">Save</button>
    <button id="cancel" class="ui button" type="button">Cancel</button>
    {{if not (eq .orderForm.Status .orderFinished .orderedCanceled)}}
        <button id="next" class="ui primary right floated button" type="button">Next Status</button>
    {{end}}
    
</form>

<script type="text/javascript">
    $(document)
        .ready(function() {
            $('.ui.form')
                .api({
                    url : "/public/orders/{{.orderForm.ID}}",
                    method : 'POST',
                    cache: false,
                    processData: false,
                    contentType: false,
                    beforeSend: (settings)=>{
                        settings.data = new FormData($(".ui.form")[0]);
                        return settings;
                    },
                    onSuccess    : function(response) {
                        $('#main_content').html(response);
                    },
                    onFailure    : function(response) {
                        $('#main_content').html(response);
                    }
                });
            $('#cancel')
                .on('click', function () {
                    $.ajax({
                        async: false,
                        type: "get",
                        url: "/public/orders",
                        success: function (data) {
                            $('#main_content').html(data);
                        }
                    });
                });
            $('#next')
                .on('click', function () {
                    $.ajax({
                        async: false,
                        type: "post",
                        url: "/public/orders/{{.orderForm.ID}}/next",
                        success: function (data) {
                            $('#main_content').html(data);
                        }
                    });
                });
            $('#cancel-order')
                .on('click', function () {
                    $.ajax({
                        async: false,
                        type: "post",
                        url: "/public/orders/{{.orderForm.ID}}/cancel",
                        success: function (data) {
                            $('#main_content').html(data);
                        }
                    });
                });
        });
</script>
