<form class="ui form" enctype="multipart/form-data">
    <h2 id="title" style="margin-top: 15px">
    <i class="shopping bag icon"></i>
    {{i18n .Lang "manager_store.title"}}
    </h2>
    {{template "public/store/form/body.tpl" .}}

    <button id="save" class="ui primary button" type="submit">{{i18n .Lang "forms.save"}}</button>
    <button id="cancel" class="ui button" type="button">{{i18n .Lang "forms.cancel"}}</button>
</form>

<form class="dropzone dz-clickable" id="myAwesomeDropzone">
    <div class="dz-message">    
        Drop files here or click to upload.
    </div>
</form>

<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/dropzone/5.2.0/min/dropzone.min.css">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/dropzone/5.2.0/min/basic.min.css">
<script src="https://cdnjs.cloudflare.com/ajax/libs/dropzone/5.2.0/min/dropzone.min.js"></script>

<script type="text/javascript">

    window.onload = function() {

        Dropzone.options.myAwesomeDropzone = {
            clickable: true,
            maxFiles: 1,
            addRemoveLinks: true,
            acceptedFiles: '.png,.jpg', 
            dictDefaultMessage: "Upload your files here",
            accept: function(file, done) {
              console.log("uploaded");
              done();
            },
            init: function() {
              this.on("maxfilesexceeded", function(file){
                  alert("No more files please!");
              });
            }
          };
    };
    $(document)
        .ready(function() {
            $('.ui.form')
                .api({
                    url : "/public/store",
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
                            type: "get",
                            url: "/public/store",
                            success: function (data) {
                                $('#main_content').html(data);
                            }
                        });
                    });
        });
            
</script>

