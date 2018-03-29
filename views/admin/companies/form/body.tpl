<div class="ui five wid raised segment">
    <div class="ui error message"></div>
    {{template "common/flash.tpl" .}}
    <div class="field">
        <div class="two fields">
            <div class="field">
                <label>{{i18n .Lang "forms.name"}}</label>
                <input name="name" value="{{.companyForm.Name}}" type="text" placeholder="{{i18n .Lang "forms.name"}}"/>
            </div>
            <div class="field">
                <label>{{i18n .Lang "forms.contact"}}</label>
                <input name="contact" type="text" value="{{.companyForm.Contact}}" placeholder="{{i18n .Lang "forms.contact"}}" />
            </div>
        </div>
    </div>
    <div class="field">
        <div class="fields">
            <div class="six wide field">
                <label>{{i18n .Lang "forms.phone_number"}}</label>
                <input name="phoneNumber" value="{{.companyForm.PhoneNumber}}" type="text" placeholder="{{i18n .Lang "forms.phone_number"}}"/>
            </div>
            <div class="ten wide field">
                <label>{{i18n .Lang "forms.email"}}</label>
                <input name="email" value="{{.companyForm.Email}}" type="email" placeholder="{{i18n .Lang "forms.email"}}"/>
            </div>
        </div>
    </div>
</div>

<script>
    $(document)
            .ready(function() {
                $('.ui.form')
                        .form({
                            fields: {
                                companyname: {
                                    identifier  : 'name',
                                    rules: [
                                        {
                                            type   : 'empty',
                                            prompt : 'Please enter your company name'
                                        }
                                    ]
                                },
                                contact: {
                                    identifier  : 'contact',
                                    rules: [
                                        {
                                            type   : 'empty',
                                            prompt : 'Please enter your contact'
                                        }
                                    ]
                                },
                                email: {
                                    identifier  : 'email',
                                    optional   : true,
                                    rules: [
                                        {
                                            type   : 'email',
                                            prompt : 'Incorrect email format'
                                        }
                                    ]
                                },
                                phone: {
                                    identifier  : 'phoneNumber',
                                    rules: [
                                        {
                                            type   : 'regExp[^[\\d+-]+$]',
                                            prompt : 'Incorrect phone number format'
                                        }
                                    ]
                                }
                            }
                        });
            });
</script>
