<h1 class="ui header" style="text-align:center;">Products</h1>
{{template "common/flash.tpl" .}}
<table class="ui single line striped collapsing table" id="table_list">
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
    {{ range .products }}
    <tr>
        <td class="center aligned">{{ .Client.Name}}</td>
        <td class="center aligned">{{ .Product.Name}}</td>
        <td class="center aligned">{{ .Status}}</td>
        <td class="center aligned">{{ .Fee}}</td>
        <td class="center aligned">
            <button type="button"
                    class="ui basic button"
                    onclick="editProducts('{{ .ID}}');">
                View
            </button>
        </td>
        <td class="center aligned">
            <button type="button"
                    class="ui negative button"
                    data-toggle="modal"
                    data-target=".bs-example-modal-sm"
                    onclick="deleteProducts('{{ .ID}}');">
                Delete
            </button>
        </td>
    </tr>
    {{ end }}
    </tbody>
</table>
</br>

<div class="ui middle aligned center aligned grid">
    <button type="button"
            title="View user"
            class="ui basic big button"
            onclick="newProducts();">
        <i class="add product icon"></i>
        Create product
    </button>
</div>

<style>
    #table_list {
       margin-left:auto;
       margin-right:auto;
     }
</style>

<script type="text/javascript">
    /*function newProducts() {
        $.ajax({
            async: false,
            type: "get",
            url: "/products/new",
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }

    function editProducts(product_id) {
        $.ajax({
            async: false,
            type: "get",
            url: "/products/" + product_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
    function deleteProducts(product_id) {
        $.ajax({
            async: false,
            type: "delete",
            url: "/products/" + product_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }*/
</script>
