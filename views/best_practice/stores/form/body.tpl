<div class="ui raised segment">
        <div class="ui error message"></div>
        {{template "best_practice/common/flash.tpl" .}}
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
            <div class="six wide field">
                <label>Manager</label>
                <div class="field">
                    <select name="manager" class="ui fluid dropdown">
                            <option value="">Manager</option>
                        {{ range .managers }}
                        {{ if $.storeForm }}
                            <option value="{{.ID}}" {{ if eq .ID $.storeForm.ManagerID}} selected {{end}}>{{.Username}}</option>
                        {{else}}
                            <option value="{{.ID}}">{{.Username}}</option>
                        {{end}}
                        {{end}}
                    </select>
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
                                    manager: {
                                        identifier  : 'manager',
                                        rules: [
                                            {
                                                type   : 'empty',
                                                prompt : 'Please select a manager'
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