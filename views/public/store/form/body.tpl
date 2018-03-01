<div class="ui raised segment">
    <div class="ui error message"></div>
    {{template "common/flash.tpl" .}}
    <div class="field">
        <div style="{{if .storeForm.Image}}
                        border: 1px solid #ddd;
                        border-radius: 4px;
                        padding: 5px;
                        width: 150px;
                    {{else}}
                        width:100%;float:left;position:relative;display:none
                    {{end}};">
            <img class="ui fluid image" id="preview" src="{{.storeForm.Image}}"/>
            <i id="imgCloser" class="close icon"></i>
            <input type="hidden" id="oldImage" name="oldImage" value="{{.storeForm.Image}}">
        </div>
        <input type="file" accept="image/*" name="image" id="poster">
    </div>

    <div class="field">
        <div class="two fields">
            <div class="field">
                <label>Name</label>
                <input name="name" value="{{.storeForm.Name}}" type="text" placeholder="Name"/>
            </div>
            <div class="ten wide field">
                <label>Address</label>
                <input name="address" value="{{.storeForm.Address}}" type="text" placeholder="Address"/>
            </div>

        </div>
    </div>
    <div class="field">
        <div class="three fields">
            <div class="field">
                <label>Phone number</label>
                <input name="phone" value="{{.storeForm.PhoneNumber}}" type="text" placeholder="Phone number"/>
            </div>
            <div class="field">
                <label>Latitude</label>
                <input name="latitude" value="{{.storeForm.Latitude}}" type="text" placeholder="Latitude"/>
            </div>
            <div class="field">
                <label>Longitude</label>
                <input name="longitude" value="{{.storeForm.Longitude}}" type="text" placeholder="Longitude"/>
            </div>
        </div>
    </div>
    <div class="two fields">
        <div class="field">
            <label>Company</label>
            <div class="field">
                <div>{{.storeForm.Company.Name}}</div>
            </div>
        </div>
        <div class="field">
            <label>Services</label>
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

</style>

<script>
    $(document)
        .ready(function() {
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


<script>
    $(document).ready( function () {
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
    });
</script>
