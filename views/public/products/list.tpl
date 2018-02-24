<h1 class="ui header" style="text-align:center;">Products</h1>
{{template "common/flash.tpl" .}}
<div class="ui divider"></div>

<table class="ui single line striped collapsing table" id="products-table"
    style="margin-left:auto; margin-right:auto; table-layout:fixed; width:100%;">
    <thead>
    <tr>
        <th class="center aligned">Name</th>
        <th class="center aligned">Description</th>
        <th class="center aligned">Value</th>
        <th class="center aligned">Image</th>
        <th class="center aligned">Service</th>
        <th class="center aligned"></th>
        <th></th>
    </tr>
    </thead>
    <tbody>
    {{ range .products }}
    <tr>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Name}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Description}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Value}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;"></td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Service.Name}}</td>
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

<div class="ui middle aligned center aligned grid">
    <button type="button"
            title="View user"
            class="ui blue basic big button"
            onclick="newProducts();"
            style="margin: 15px;">
        <i class="add product icon"></i>
        Create product
    </button>
</div>

<script type="text/javascript">

    $(document).ready(function() {
        $('#products-table').DataTable();
    });


    function newProducts() {
        $.ajax({
            async: false,
            type: "get",
            url: "/public/products/new",
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }

    function editProducts(product_id) {
        $.ajax({
            async: false,
            type: "get",
            url: "/public/products/" + product_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }

    function deleteProducts(product_id) {
        $.ajax({
            async: false,
            type: "delete",
            url: "/public/products/" + product_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
</script>
