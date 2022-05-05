import React, { useState } from "react";
import { FlatList, ListRenderItem, Text, View } from "react-native";
import Item from "../models/item";
import List from "../models/list";
import ItemComponent from "./item_component";

const items: Item[] = [
    {
        id: "item_id1",
        task: "task1",
        priority: 0,
        list_id: "list_id1"
    },
    {
        id: "item_id2",
        task: "task2",
        priority: 1,
        list_id: "list_id2"
    },
    {
        id: "item_id3",
        task: "random task",
        priority: 1,
        list_id: "list_id3"
    }
];

interface ItemListProps {
    list: List;
}

const ItemList: React.FC<ItemListProps> = ({ list }) => {
    const [items, setItems] = useState<Item[]>([]);

    const renderItem: ListRenderItem<Item> = ({ item }) => (
        <ItemComponent item={item}/>
    );

    return (
        <View>
            <Text>{list.name}</Text>
            <FlatList data={items} renderItem={renderItem} keyExtractor={item => item.id}/>
        </View>
    );
}

export default ItemList;