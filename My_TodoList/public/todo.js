(function($) {
    'use strict';
    $(function() {
    var todoListItem = $('.todo-list');
    var todoListInput = $('.todo-list-input');
    $('.todo-list-add-btn').on("click", function(event) { //버튼 눌리면 실행될 내용
        event.preventDefault();
        
        var item = $(this).prevAll('.todo-list-input').val();
        //item과 id, created_at 등의 내용을 json으로 바꿔서 특정 url로 전송

        if (item) {
            $.post("/TodoList",{name:item},addItem) //Jsond
            //todoListItem.append("<li><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
            todoListInput.val("");
        }
    
    });

    var addItem = function(item){
        console.log(item)
        if (item.completed){
            todoListItem.append("<li class='completed'"+ " id='" + item.ID + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' checked='checked'/>" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");    
        }
        else{
            todoListItem.append("<li "+ " id='" + item.ID + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
        }
        todoListInput.val("");
    }

    $.get("/TodoList",function(items){
        items.forEach(element => {
            addItem(element)
        });
    })



    todoListItem.on('change', '.checkbox', function() {
        var id = $(self).closest("li").attr('id')
        var $self = $(this);
        
        var complete = true;
        if ($(self).attr('checked')){
            complete = false
        }
        $.get("complete-todo/"+id, function(data){

            if (data.success == true){
                if (data.complition == true){
                    $(self).removeAttr('checked');
                }else{
                    $(self).attr('checked', 'checked');
                }
            }

            // if ($(self).attr('checked')) {
            //     $(self).removeAttr('checked');
            // } else {
            //     $(self).attr('checked', 'checked');
            // }
            
            $(self).closest("li").toggleClass('completed');
        })
        
        

        
        
    });
        

    todoListItem.on('click', '.remove', function() {
            var id = $(this).closest("li").attr('id')
            var $self = $(this);
            $.ajax({
                url:"TodoList/" + id,
                type: "DELETE",
                success: function(data){ //응답
                    if(data.success){
                        $self.parent().remove();
                    }else{
                        console.log("fail to delete")
                    }
                }
            })
        });
        
    });
})(jQuery);