import React from "react";
import ItemList from "../components/item_list";

const ItemScreen = () => {
    return (
        <ItemList list={{id: "unique_list1", name: "list1", user_id: "user1"}} />
    );
}

export default ItemScreen;