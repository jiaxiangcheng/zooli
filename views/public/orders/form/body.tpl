<div class="ui five wid raised segment">
    <div class="ui error message">
        {{template "common/flash.tpl" .}}
    </div>
    <div class="field">
        <div class="three fields">
            <div class="four wide field">
                <label>Client</label>
                <div class="ui disabled input">
                    <input value="{{.orderForm.Client.Name}}" type="text"/>
                </div>

            </div>
            <div class="six wide field">
                <label>Product</label>
                <div class="field">
                    <div class="ui disabled input">
                        <input type="text" value="{{.orderForm.Product.Name}}">
                    </div>
                </div>
            </div>
            <div class="four wide field">
                <label>Status</label>
                <a class="ui {{if eq .orderForm.Status .ordered}}orange{{else if eq .orderForm.Status .inService}}yellow{{else if eq .orderForm.Status .endService}}olive{{else if eq .orderForm.Status .waitingForPayment}}teal{{else if eq .orderForm.Status .orderFinished}}green{{else}}grey{{end}} label">{{.orderForm.Status}}</a>
            </div>
        </div>
        <div class="field">
            <div class="ui right icon input">
                <i class="yen sign icon"></i>
                <input type="text" placeholder="Amount" value="{{.orderForm.Fee}}">
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
                        value: {
                            identifier  : 'fee',
                            optional   : true,
                            rules: [
                                {
                                    type   : 'number',
                                    prompt : 'Incorrect order value format'
                                }
                            ]
                        }
                    }
                });
    });

</script>