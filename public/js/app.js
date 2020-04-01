"use strict";
var idPrefix = "#table_id_";
var urlPrefix = "#table_url_";
var inputPostfix = "_input";
var actionsPrefix = "#table_actions_";

function openEditRow() {
    var id = this.parentNode.id;
    var idId = idPrefix + id;
    var urlId = urlPrefix + id;

    var oldId = $(idId).text();
    var oldUrl = $(urlId).text();

    $(idId).data("old", oldId);
    $(idId).html(toInput(idId, oldId));
    $(urlId).html(toInput(urlId, oldUrl));

    $(actionsPrefix + id).html(cancelOrSaveButtons(id))
}

function cancelEditRow() {
    var id = this.parentNode.id;

    var idVal = $(idPrefix + id + inputPostfix).val();
    var urlVal = $(urlPrefix + id + inputPostfix).val();

    $(idPrefix + id).html("<a href='/l/" + idVal + "'>" + idVal + "</a>");
    $(urlPrefix + id).html("<a href='" + urlVal + "'>" + urlVal + "</a>");
    $(actionsPrefix + id).html(editOrDeleteButtons(id));
}

function saveEditLink() {
    var id = this.parentNode.id;

    var oldId = $(idPrefix + id).data("old");
    var idVal = $(idPrefix + id + inputPostfix).val();
    var urlVal = $(urlPrefix + id + inputPostfix).val();

    if (!urlVal || !idVal) {
        return alert("Empty Value(s)");
    }

    if (!oldId) {
        return alert("Old Id not found!!!!");
    }

    $.ajax({
        url: "/v1/link/" + oldId,
        data: JSON.stringify({id: idVal, url: urlVal}),
        contentType: "application/json",
        type: "PUT",
        success: function success() {
            location.href = "/v1/";
        },
        error: function error(data, status, err) {
            alert(err);
        }
    });
}

function saveLink() {
    var newId = $("#new_id").val();
    var newUrl = $("#new_url").val();
    if (!newUrl) {
        return alert("Empty Link");
    }
    $.ajax({
        url: "/v1/link",
        data: JSON.stringify({id: newId, url: newUrl}),
        contentType: "application/json",
        type: "POST",
        success: function success() {
            location.href = "/v1/";
        },
        error: function error(data, status, err) {
            return alert(err);
        }
    });
}

function openDeleteModal() {
    var id = this.parentNode.id;
    $('#delete_button').data("id", id);
    $('.basic.modal').modal("setting", "closeable", false).modal("show");
}

function deleteRow() {
    var id = $('#delete_button').data("id");
    $.ajax({
        url: "/v1/link/" + id,
        type: "DELETE",
        success: function success() {
            location.href = "/v1/";
        },
        error: function error(data, status, err) {
            alert(err);
        }
    })
}

function toInput(id, val) {
    return '<div class="ui input small fluid"><input id="' + id.substr(1) + inputPostfix + '" type="text" value="' + val + '"/></div>';
}

function cancelOrSaveButtons(id) {
    return '<div class="ui buttons" id="' + id + '">' +
        '<button class="ui green inverted small button" onclick="saveEditLink.call(this)"><i class="check icon"></i>Save</button>' +
        '<button class="ui red inverted small button" onclick="cancelEditRow.call(this)"><i class="close icon"></i>Cancel</button>' +
        '</div>';
}

function editOrDeleteButtons(id) {
    return '<div class="ui buttons"  id="' + id + '">' +
        '<button class="ui icon teal mini button clip" onclick="copyToClipboard.call(this)"><i class="copy icon"></i></button>' +
        '<button class="ui icon olive mini button edit" onclick="openEditRow.call(this)"><i class="pen icon"></i></button>' +
        '<button class="ui icon red mini button delete" onclick="openDeleteModal.call(this)"><i class="trash icon"></i></button>' +
        '</div>';
}

function copyToClipboard() {
    var id = this.parentNode.id;
    var textarea = document.createElement("textarea");
    textarea.textContent = location.origin + "/l/" + id;
    textarea.style.position = "fixed"; // Prevent scrolling to bottom of page in MS Edge.
    document.body.appendChild(textarea);
    textarea.select();
    document.execCommand("copy");

    document.body.removeChild(textarea);
    $("body").toast({
        class: "inverted",
        position: "bottom right",
        message: "Link Copied to Clipboard"
    });
}

$(function () {
    // Format Dates to Locale Timezone & Format
    $(".created-at").each(function () {
        $(this).text(new Date(this.innerText).toLocaleString());
    });

    $("#logout").on('click', function () {
        // Invalidate session cookie
        document.cookie = "session-status=invalid; path=/";
        // Reload page
        location.href = "/login";
    })
});