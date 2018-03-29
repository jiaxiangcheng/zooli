<div class="ui five wid raised segment">
    <div class="ui error message"></div>
    {{template "common/flash.tpl" .}}
    <div class="field">
        <div class="three fields">
            <div class="four wide field">
                <label>{{i18n .Lang "form.name"}}</label>
                <input name="name" value="{{.productForm.Name}}" type="text" placeholder="{{i18n .Lang "form.name"}}"/>
            </div>
            <div class="six wide field">
                <label>{{i18n .Lang "form.value"}}</label>
                <input name="value" value="{{.productForm.Value}}" type="text" placeholder="{{i18n .Lang "form.value"}}"/>
            </div>
            <div class="six wide field">
                <label>{{i18n .Lang "form.services"}}</label>
                <div class="field">
                    <select name="service" class="ui fluid dropdown">
                        <option value="">{{i18n .Lang "form.services"}}</option>
                    {{ range .services }}
                    {{ if $.productForm }}
                        <option value="{{.ID}}" {{ if eq .ID $.productForm.ServiceID}} selected {{end}}>{{.Name}}</option>
                    {{else}}
                        <option value="{{.ID}}">{{.Name}}</option>
                    {{end}}
                    {{end}}
                    </select>
                </div>
            </div>
        </div>
    </div>
    <div class="field">
        <label>{{i18n .Lang "form.description"}}</label>
        <textarea name="description" type="text" placeholder="{{i18n .Lang "form.type_description"}}">{{.productForm.Description}}</textarea>
    </div>
    <div class="field" id="image-container">
        <div style="{{if .productForm.Images}}
                        border: 1px solid #ddd;
                        border-radius: 4px;
                        padding: 5px;
                        width: 150px;
                    {{else}}
                        width:100%;float:left;position:relative;display:none
                    {{end}};">
            <img class="ui fluid image" id="preview" src="{{.productForm.Images}}"/>
            <i id="imgCloser" class="close icon" style="position: absolute;top:15px;right:15px;cursor: pointer;"></i>
            <input type="hidden" id="oldImage" name="oldImage" value="{{.productForm.Images}}">
        </div>
        <input type="file" accept="image/*" name="image" id="poster">
    </div>
</div>

<div id="image_modal" class="ui modal">
    <div class="image content">
        <img id="modal-image">
    </div>
</div>
    
<style>
    #image-container{
        position: relative;
    }

    #imgCloser{
        position: absolute;
        top: 0;
        right: -10;
        cursor: pointer;
    }

    img:hover{
        cursor: pointer;
        transition: 0.3s;
        opacity: 0.5
    }

</style>

<script>
    $(document)
            .ready(function() {
                var preview = document.getElementById('preview');
                $("#imgCloser").click(function () {
                    $('#poster').val('');
                    $('#oldImage').val("");
                    preview.src = "";
                    preview.parentNode.style = 'width:100%;float:left;position:relative;display:none';
                });

                $("#poster").change(function () {
                    if (event.target.files.length > 0) {
                        preview.src = URL.createObjectURL(event.target.files[0]);
                        preview.parentNode.style = 'border: 1px solid #ddd;border-radius: 4px;padding: 5px;width: 150px;';
                    } else {
                        preview.src = "";
                        preview.parentNode.style = 'width:100%;float:left;position:relative;display:none';
                    }
                });

                $("#preview").click(function () {
                    $("#modal-image").attr("src", preview.src);
                    $('#image_modal').modal('show');
                });

                $('.dropdown').dropdown();
                $('.ui.form')
                        .form({
                            fields: {
                                name: {
                                    identifier  : 'name',
                                    rules: [
                                        {
                                            type   : 'empty',
                                            prompt : 'Please enter your product name'
                                        }
                                    ]
                                },
                                description: {
                                    identifier  : 'description',
                                    rules: [
                                        {
                                            type   : 'empty',
                                            prompt : 'Please enter your description'
                                        }
                                    ]
                                },
                                value: {
                                    identifier  : 'value',
                                    optional   : true,
                                    rules: [
                                        {
                                            type   : 'number',
                                            prompt : 'Incorrect product value format'
                                        }
                                    ]
                                },
                                service: {
                                    identifier  : 'service',
                                    rules: [
                                        {
                                            type   : 'empty',
                                            prompt : 'Incorrect service name format'
                                        }
                                    ]
                                }
                            }
                        });
            });
</script>
