<div class="ui five wid raised segment">
    <div class="ui error message"></div>
    {{template "best_practice/common/flash.tpl" .}}
    <div class="field">
        <div class="two fields">
            <div class="field">
                <label>Name</label>
                <input name="name" value="{{.companyForm.Name}}" type="text" placeholder="Name"/>
            </div>
            <div class="field">
                <label>Contact</label>
                <input name="contact" type="text" value="{{.companyForm.Contact}}" placeholder="Contact" />
            </div>
        </div>
    </div>
    <div class="field">
        <div class="fields">
            <div class="six wide field">
                <label>Phone number</label>
                <input name="phoneNumber" value="{{.companyForm.PhoneNumber}}" type="text" placeholder="Phone number"/>
            </div>
            <div class="ten wide field">
                <label>Email</label>
                <input name="email" value="{{.companyForm.Email}}" type="email" placeholder="Email"/>
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
                                            type   : 'number',
                                            prompt : 'Incorrect phone number format'
                                        }
                                    ]
                                }
                            }
                        });
            });
</script>
