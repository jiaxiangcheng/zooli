<div class="ui five wid raised segment">
    <div class="ui error message"></div>
    {{template "common/flash.tpl" .}}
    <div class="field">
        <div class="three fields">
            <div class="four wide field">
                <label>Name</label>
                <input name="name" value="{{.productForm.Name}}" type="text" placeholder="Name"/>
            </div>
            <div class="six wide field">
                <label>Value</label>
                <input name="value" value="{{.productForm.Value}}" type="text" placeholder="Value"/>
            </div>
            <div class="six wide field">
                <label>Services</label>
                <div class="field">
                        <div class="ui multiple selection dropdown" id="services">
                            <!-- This will receive comma separated value like 1,2,3 !-->
                            <input name="services" type="hidden" >
                            <i class="dropdown icon"></i>
                            <div class="default text">Services</div>
                            <div class="menu">
                                {{ range .services }}
                                    <div class="item" data-value="{{.Name}}">{{.Name}}</div>
                                {{end}}
                            </div>
                        </div>
                </div>
            </div>
        </div>
    </div>
    <div class="field">
        <label>Description</label>
        <textarea name="description"
        type="text" value="{{.productForm.Description}}"
        placeholder="Type your product description"></textarea>
    </div>
    <div class="field">
        <div style="width:100%;float:left;position:relative;display:{{if .storeForm.Image}}inline{{else}}none{{end}};">
            <img class="ui fluid image" id="preview" src="{{.storeForm.Image}}" style="width:100%;max-height:100%;"/>
            <i id="imgCloser" class="close icon" style="position: absolute;top:15px;right:15px;cursor: pointer;"></i>
            <input type="hidden" id="oldImage" name="oldImage" value="{{.storeForm.Image}}">
        </div>
        <input type="file" accept="image/*" name="image" id="poster">
    </div>
</div>

<script>
    $(document)
            .ready(function() {
                $('.dropdown').dropdown();
                {{if .productForm}}
                    $('#services').dropdown('set selected', [{{range $i, $s := .storeForm.Services}}{{if $i}},{{end}}{{$s.Name}}{{end}}]);
                {{end}}
                $('.ui.form')
                        .form({
                            fields: {
                                productname: {
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
                                    rules: [
                                        {
                                            type   : 'number',
                                            prompt : 'Incorrect product value format'
                                        }
                                    ]
                                },
                                image: {
                                    identifier  : 'phoneNumber',
                                    rules: [
                                        {
                                            type   : 'number',
                                            prompt : 'Incorrect phone number format'
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


<script>
    $(document).ready( function () {
        var preview = document.getElementById('preview');
        $("#imgCloser").click(function () {
            $('#poster').val('');
            $('#oldImage').val("");
            preview.src = "";
            preview.parentNode.style.display = 'none';
        });

        $("#poster").change(function () {
            if (event.target.files.length > 0) {
                preview.src = URL.createObjectURL(event.target.files[0]);
                preview.parentNode.style.display = 'inline';
            } else {
                preview.src = "";
                preview.parentNode.style.display = 'none';
            }
        });



    });
</script>
