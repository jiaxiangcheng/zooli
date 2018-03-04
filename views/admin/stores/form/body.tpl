<div class="ui raised segment">
    <div class="ui error message"></div>
    {{template "common/flash.tpl" .}}
    
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
            <label>{{i18n .Lang "forms.companies"}}</label>
            <div class="field">
                <select name="company" class="ui fluid dropdown">
                    <option value="">{{i18n .Lang "forms.companies"}}</option>
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
            <label>{{i18n .Lang "forms.services"}}</label>
            <div class="field">
                <div class="ui multiple selection dropdown" id="services">
                    <!-- This will receive comma separated value like 1,2,3 !-->
                    <input name="services" type="hidden" >
                    <i class="dropdown icon"></i>
                    <div class="default text">{{i18n .Lang "forms.services"}}</div>
                    <div class="menu">
                    {{ range .services }}
                        <div class="item" data-value="{{.Name}}">{{.Name}}</div>
                    {{end}}
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="field">
        <label>{{i18n .Lang "forms.mangers"}}</label>
        <div class="ui fluid multiple selection dropdown" id="managers">
            <input name="managers" type="hidden">
            <i class="dropdown icon"></i>
            <span class="default text">{{i18n .Lang "forms.mangers"}}</span>
            <div class="menu">
                <div class="ui icon search input">
                    <i class="search icon"></i>
                    <input type="text" placeholder="{{i18n .Lang "forms.search_managers"}}...">
                </div>
                <div class="scrolling menu">
                    {{range $.managers}}
                        <div class="item" data-value="{{.ID}}">
                            <span class="text">{{.Username}} ({{.Name}})</span>
                        </div>
                    {{end}}
                </div>
            </div>
        </div>
    </div>
</div>

<script>
    $(document)
            .ready(function() {

                $('.dropdown').dropdown();
                {{if .storeForm}}
                    $('#services').dropdown('set selected', [{{range $i, $s := .storeForm.Services}}{{if $i}},{{end}}{{$s.Name}}{{end}}]);
                {{end}}
                {{if .storeForm}}
                    $('#managers').dropdown('set selected', [{{range $i, $s := .storeForm.Managers}}{{if $i}},{{end}}"{{$s.ID}}"{{end}}]);
                {{end}}

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
