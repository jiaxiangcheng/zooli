<div class="ui raised segment">
    <div class="ui error message"></div>
    {{template "common/flash.tpl" .}}
    <div class="field" id="image-container">
        <input type="file" id="files" name="files[]" accept="image/*" multiple/>
        <output id="store-images"></output>
    </div>

    <div class="field">
        <div class="two fields">
            <div class="field">
                <label>{{i18n .Lang "forms.name"}}</label>
                <input name="name" value="{{.storeForm.Name}}" type="text" placeholder="{{i18n .Lang "forms.name"}}"/>
            </div>
            <div class="ten wide field">
                <label>{{i18n .Lang "forms.address"}}</label>
                <input name="address" value="{{.storeForm.Address}}" type="text" placeholder="{{i18n .Lang "forms.address"}}"/>
            </div>

        </div>
    </div>
    <div class="field">
        <div class="three fields">
            <div class="field">
                <label>{{i18n .Lang "forms.phone_number"}}</label>
                <input name="phone" value="{{.storeForm.PhoneNumber}}" type="text" placeholder="{{i18n .Lang "forms.phone_number"}}"/>
            </div>
            <div class="field">
                <label>{{i18n .Lang "forms.latitude"}}</label>
                <input name="latitude" value="{{.storeForm.Latitude}}" type="text" placeholder="{{i18n .Lang "forms.latitude"}}"/>
            </div>
            <div class="field">
                <label>{{i18n .Lang "forms.longitude"}}</label>
                <input name="longitude" value="{{.storeForm.Longitude}}" type="text" placeholder="{{i18n .Lang "forms.longitude"}}"/>
            </div>
        </div>
    </div>
    <div class="two fields">
        <div class="field">
            <label>{{i18n .Lang "forms.company"}}</label>
            <div class="field">
                <div>{{.storeForm.Company.Name}}</div>
            </div>
        </div>
        <div class="field">
            <label>{{i18n .Lang "forms.services"}}</label>
            <div class="field">
                <div>{{range .storeForm.Services}} <a class="ui blue label">{{.Name}}</a> {{end}}</div>
            </div>
        </div>
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

    .thumbnail {
        border: 1px solid #ddd;
        border-radius: 4px;
        padding: 5px;
        width: 150px;
    }
</style>

<script>

    function handleImageSelectOrDragDrop(evt) {
        var files = evt.target.files;
        for (var i = 0, file; file = files[i]; i++) {
            
            if (!file.type.match('image.*')) {
                continue;
            }

            const objectURL = window.URL.createObjectURL(file);
            console.log(objectURL);
            var reader = new FileReader();
            reader.onload = (function(theFile) {
                return function(e) {
                    var span = document.createElement('span');
                    span.innerHTML = ['<a href="', objectURL, '" data-lightbox="roadtrip" data-title="', escape(theFile.name), '"><img class="thumbnail" src="', objectURL, '"/></a>'].join('');
                    document.getElementById('store-images').insertBefore(span, null);
                };
            })(file);

            reader.readAsDataURL(file);
        };
    }

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

        document.getElementById('files').addEventListener('change', handleImageSelectOrDragDrop, false);

        $('.dropdown').dropdown();

        $('.ui.form')
            .form({
                fields: {
                    name: {
                        identifier  : 'name',
                        rules: [
                            {
                                type   : 'empty',
                                prompt : 'Please enter your store name'
                            }
                        ]
                    },
                    address: {
                        identifier  : 'address',
                        rules: [
                            {
                                type   : 'empty',
                                prompt : 'Please enter your address'
                            }
                        ]
                    },
                    phone: {
                        identifier  : 'phone',
                        rules: [
                            {
                                type   : 'regExp[^[\\d+-]+$]',
                                prompt : 'Incorrect phone number format'
                            }
                        ]
                    },
                    company: {
                        identifier  : 'company',
                        rules: [
                            {
                                type   : 'empty',
                                prompt : 'Please select a company'
                            }
                        ]
                    },
                    latitude: {
                        identifier  : 'latitude',
                        rules: [
                            {
                                type   : 'empty',
                                prompt : 'Please enter a latitude value'
                            }
                        ]
                    },
                    longitude: {
                        identifier  : 'longitude',
                        rules: [
                            {
                                type   : 'empty',
                                prompt : 'Please enter a longitude value'
                            }
                        ]
                    }
                }
            });
        });
</script>

