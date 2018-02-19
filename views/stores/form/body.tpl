<div class="ui raised segment">
    <div class="ui error message"></div>
{{template "common/flash.tpl" .}}


    <div class="field">
        <div style="width:100%;float:left;position:relative;display: none;">
            <img class="ui fluid image" id="preview" src="" style="width:100%;max-height:100%;"/>
            <i id="imgCloser" class="close icon" style="position: absolute;top:15px;right:15px;cursor: pointer;"></i>
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
        <div class="six wide field">
            <label>Company</label>
            <div class="field">
                <select name="company" class="ui fluid dropdown">
                    <option value="">Company</option>
                {{ range .companies }}
                {{ if $.storeForm }}
                    <option value="{{.ID}}" {{ if eq .ID $.storeForm.CompanyID}} selected {{end}}>{{.Name}}</option>
                {{else}}
                    <option value="{{.ID}}">{{.Name}}</option>
                {{end}}
                {{end}}
                </select>
            </div>
        </div>
        <div class="field">
            <label>Services</label>
            <div class="field">
                <div class="ui multiple selection dropdown">
                    <!-- This will receive comma separated value like 1,2,3 !-->
                    <input name="services" type="hidden">
                    <i class="dropdown icon"></i>
                    <div class="default text">Services</div>
                    <div class="menu">
                    {{ range .services }}
                        <div class="item" data-value="{{.ID}}">{{.Name}}</div>
                    {{end}}
                    </div>
                </div>
            </div>
        </div>


    </div>
</div>


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
                                            type   : 'number',
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

        if ({{.storeForm.Image}}) {
            preview.src = {{.storeForm.Image}};
            preview.parentNode.style.display = 'inline';
            $('#oldImage').val({{.StoreForm.Image}});
        }
        $("#imgcloser").click(function () {
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