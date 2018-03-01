<h1 class="ui header" style="text-align:center; margin-top: 20px;">Orders</h1>
{{template "common/modal.tpl" .}}
{{template "common/flash.tpl" .}}
<div class="ui divider"></div>

<table class="ui single line striped collapsing table" id="orders-table"
    style="margin-left:auto; margin-right:auto; table-layout:fixed; width:100%;">
    <thead>
    <tr>
        <th class="center aligned">Customer</th>
        <th class="center aligned">Product</th>
        <th class="center aligned">Status</th>
        <th class="center aligned">Fee</th>
        <th class="center aligned"></th>
        <th></th>
    </tr>
    </thead>
    <tbody>
    {{ range .orders }}
    <tr>
        <td class="center aligned">{{ .Customer.Name}}</td>
        <td class="center aligned">{{ .Product.Name}}</td>
        <td class="center aligned"><a class="ui {{if eq .Status $.ordered}}orange{{else if eq .Status $.inService}}yellow{{else if eq .Status $.endService}}olive{{else if eq .Status $.waitingForPayment}}teal{{else if eq .Status $.orderFinished}}green{{else}}grey{{end}} label">{{.Status}}</a></td>
        <td class="center aligned">{{ .Fee}}</td>
        <td class="center aligned">
            <button type="button"
                    class="ui basic button"
                    onclick="editOrder('{{ .ID}}');">
                View
            </button>
        </td>
        <td class="center aligned">
            <button type="button"
                    class="ui negative button"
                    data-toggle="modal"
                    data-target=".bs-example-modal-sm"
                    onclick="deleteOrder('{{ .ID}}');">
                Delete
            </button>
        </td>
    </tr>
    {{ end }}
    </tbody>
</table>

<div class="ui middle aligned center aligned grid">
    <button type="button"
            title="View user"
            class="ui blue basic big button"
            onclick="newOrder();"
            style="margin: 15px;">
        <i class="add order icon"></i>
        Create order
    </button>
</div>

<style>
    #table_list {
       margin-left:auto;
       margin-right:auto;
     }
</style>

<script type="text/javascript">

    $(document).ready(function() {
        $('#orders-table').DataTable();
    });

    function newOrder() {
        $.ajax({
            async: false,
            type: "get",
            url: "/public/orders/new",
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }

    function editOrder(order_id) {
        $.ajax({
            async: false,
            type: "get",
            url: "/public/orders/" + order_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }

    function deleteOrder(order_id) {
        $.ajax({
            async: false,
            type: "delete",
            url: "/public/orders/" + order_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
</script>
