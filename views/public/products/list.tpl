{{template "common/modal.tpl" .}}

<div class="row">
    <div class="column">

        <div class="ui segments">
            <div class="ui segment">
                <h1 class="ui header center aligned">{{i18n .Lang "products_table.title"}}</h1>
            </div>
            <div class="ui segment">
                {{template "common/flash.tpl" .}}
                <table class="ui compact fixed single line selectable striped celled table tablet stackable" id="data_table" cellspacing="0" width="100%">
                    <thead>
                        <tr>
                            <th>{{i18n .Lang "table_attribute_names.name"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.description"}}</th>
                            <th class="center aligned">{{i18n .Lang "table_attribute_names.value"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.services"}}</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range .products }}
                        <tr>
                            <td>{{ .Name}}</td>
                            <td>{{ .Description}}</td>
                            <td class="center aligned">{{ .Value}}</td>
                            <td>{{ .Service.Name}}</td>
                            <td class="center aligned">
                                <i class="blue link pencil alternate icon" onclick="editProduct('{{ .ID}}');"></i>
                                <i class="red link trash alternate icon" onclick="openDeleteModal('{{ .ID}}');"></i>
                            </td>
                        </tr>
                        {{ end }}
                    </tbody>
                    <tfoot class="full-width">
                        <tr>
                            <th colspan="5">
                                <div class="ui right floated small primary labeled icon button" onclick="newProduct();">
                                    <i class="cubes icon"></i> {{i18n .Lang "products_table.add_product"}}
                                </div>
                            </th>
                        </tr>
                    </tfoot>
                </table>
            </div>
        </div>
    </div>
</div>


<script type="text/javascript">

$(document).ready(function() {
    $('#data_table').DataTable({
        language: {
            "search": {{i18n .Lang "search input"}}
        },
        lengthChange: false,
        info: false
    }
);

});

function newProduct() {
    $.ajax({
        async: false,
        type: "get",
        url: "/public/products/new",
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}

function editProduct(product_id) {
    $.ajax({
        async: false,
        type: "get",
        url: "/public/products/" + product_id,
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}

function deleteProduct(product_id) {
    $.ajax({
        async: false,
        type: "delete",
        url: "/public/products/" + product_id,
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}
function openDeleteModal(product_id) {
    $('#mini_modal .header').html("Alert");
    $('#mini_modal .content').html("Are you sure to delete product?");
    $('#mini_modal')
    .modal({
        onApprove : function() {
            deleteProduct(product_id)
        }
    })
    .modal('show');
}
</script>
