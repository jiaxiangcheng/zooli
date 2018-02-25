<div class="ui five wid raised segment">
    <div class="ui error message">
        {{template "common/flash.tpl" .}}
    </div>
    <div class="field">
        <div class="three fields">
            <div class="four wide field">
                <label>Client</label>
                <input name="name" value="{{.orderForm.Client.Name}}" type="text" placeholder="Client Name"/>
            </div>
            <div class="six wide field">
                <label>Product</label>
                <div class="field">
                    <select name="product" class="ui fluid dropdown">
                        <option value="">Products</option>
                        {{ range .products }}
                            {{ if $.orderForm }}
                                <option value="{{.ID}}" {{ if eq .ID $.orderForm.ProductID}} selected {{end}}>{{.Name}}</option>
                            {{else}}
                                <option value="{{.ID}}">{{.Name}}</option>
                            {{end}}
                        {{end}}
                    </select>
                </div>
            </div>
            <div class="four wide field">
                <label>Status</label>
                <div name="status" value="{{.orderForm.Status}}">Ordered</div>
            </div>
        </div>
        <div class="field">
            <label>Fee</label>
            <div class="field">
                <i class="yen sign icon"></i><input name="fee" value="{{.orderForm.Fee}}" type="text" placeholder="Order Fee Value"/>
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
                                    prompt : 'Please enter your product name'
                                }
                            ]
                        },
                        value: {
                            identifier  : 'fee',
                            optional   : true,
                            rules: [
                                {
                                    type   : 'number',
                                    prompt : 'Incorrect order value format'
                                }
                            ]
                        },
                        service: {
                            identifier  : 'product',
                            rules: [
                                {
                                    type   : 'empty',
                                    prompt : 'Incorrect product name format'
                                }
                            ]
                        }
                    }
                });
    });

</script>