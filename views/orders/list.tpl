<h1 class="ui header" style="text-align:center;">Orders</h1>
{{template "common/flash.tpl" .}}
<table class="ui single line striped collapsing table" id="table_list" style="table-layout:fixed; width:100%;">
    <thead>
    <tr>
        <th class="center aligned">Client</th>
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
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Client.Name}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Product.Name}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Status}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Fee}}</td>
        <td class="center aligned">
            <button type="button"
                    class="ui basic button"
                    onclick="editOrders('{{ .ID}}');">
                View
            </button>
        </td>
        <td class="center aligned">
            <button type="button"
                    class="ui negative button"
                    data-toggle="modal"
                    data-target=".bs-example-modal-sm"
                    onclick="deleteOrders('{{ .ID}}');">
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
            class="ui basic big button"
            onclick="newOrders();"
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
    /*function newOrders() {
        $.ajax({
            async: false,
            type: "get",
            url: "/orders/new",
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }

    function editOrders(order_id) {
        $.ajax({
            async: false,
            type: "get",
            url: "/orders/" + order_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
    function deleteOrders(order_id) {
        $.ajax({
            async: false,
            type: "delete",
            url: "/orders/" + order_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }*/
</script>
