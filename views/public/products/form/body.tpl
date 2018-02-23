<div class="ui five wid raised segment">
    <div class="ui error message"></div>
    {{template "common/flash.tpl" .}}
    <div class="field">
        <div class="two fields">
            <div class="field">
                <label>Name</label>
                <input name="name" value="{{.productForm.Name}}" type="text" placeholder="Name"/>
            </div>
            <div class="field">
                <label>Description</label>
                <input name="description" type="text" value="{{.productForm.Description}}" placeholder="Description" />
            </div>
        </div>
    </div>
    <div class="field">
        <div class="fields">
            <div class="four wide field">
                <label>Value</label>
                <input name="value" value="{{.productForm.Value}}" type="text" placeholder="Value"/>
            </div>
            <div class="four wide field">
                <label>Image</label>
                <input name="image" value="{{.productForm.Image}}" type="text" placeholder="Image"/>
            </div>
            <div class="eight wide field">
                <label>Service</label>
                <input name="service" value="{{.productForm.Service}}" type="text" placeholder="Service"/>
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
