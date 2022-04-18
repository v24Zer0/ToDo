import React from "react";
import { FlatList, ListRenderItem, Text, View } from "react-native";
import Item from "../models/item";
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

const ItemList = () => {
    const renderItem: ListRenderItem<Item> = ({ item }) => (
        <ItemComponent id={item.id} task={item.task} priority={item.priority} list_id={item.list_id} />
    );

    return (
        <View>
            <Text>
                Hello World!
            </Text>
            <FlatList data={items} renderItem={renderItem} keyExtractor={item => item.id}/>
        </View>
    );
}

export default ItemList;