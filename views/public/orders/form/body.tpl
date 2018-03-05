<div class="ui five wid raised segment">
    <div class="ui error message">
        {{template "common/flash.tpl" .}}
    </div>
    <div class="field">
        <div class="three fields">
            <div class="four wide field">
                <label>Customer</label>
                <div class="ui disabled input">
                    <input value="{{.orderForm.Customer.Name}}" type="text"/>
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
                <label>Current Status</label>
                <a class="ui {{if eq .orderForm.Status .ordered}}orange{{else if eq .orderForm.Status .inService}}yellow{{else if eq .orderForm.Status .endService}}olive{{else if eq .orderForm.Status .waitingForPayment}}teal{{else if eq .orderForm.Status .orderFinished}}green{{else}}grey{{end}} label">{{.orderForm.Status}}</a>
            </div>
            <div class="four wide field">
                {{if not (eq .orderForm.Status .orderedCanceled)}}
                    <button id="cancel-order" class="ui red right floated button">Cancel order</button>
                {{end}}
            </div>
        </div>

        <div class="ui mini steps">
            {{ range .orderLogs }}
                {{ if eq $.orderForm.Status .Status}}
                    <a class="active step">
                        <div class="content">
                            <div class="title">{{$.orderForm.Status}}</div>
                            <div class="description">{{if eq $.orderForm.Status $.ordered}}In order{{else if eq $.orderForm.Status $.inService}}In service{{else if eq $.orderForm.Status $.endService}}End of service{{else if eq $.orderForm.Status $.waitingForPayment}}Waiting for payment{{else if eq $.orderForm.Status $.orderFinished}}Order finished{{else}}Order canceled{{end}}</div>
                            <div class="description">{{.Timestamp}}</div>
                        </div>
                    </a>
                {{else}}
                    <a class="step">
                        <div class="content">
                            <div class="title">{{.Status}}</div>
                            <div class="description">{{if eq .Status $.ordered}}In order{{else if eq .Status $.inService}}In service{{else if eq .Status $.endService}}End of service{{else if eq .Status $.waitingForPayment}}Waiting for payment{{else if eq .Status $.orderFinished}}Order finished{{else}}Order canceled{{end}}</div>
                            <div class="description">{{.Timestamp}}</div>
                        </div>
                    </a>
                {{end}}
            {{end}}
            {{ if .nextStep }}
                <div class="disabled step">
                    <div class="content">
                      <div class="title">{{.nextStep}}</div>
                    </div>
                </div>
            {{end}}
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