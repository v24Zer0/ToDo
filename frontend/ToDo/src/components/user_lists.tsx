import React from "react";
import { FlatList, ListRenderItem } from "react-native";
import List from "../models/list";
import ListComponent from "./list_component";

const listData: List[] = [
    {
        id: "unique_list1",
        name: "list1",
        user_id: "user1"
    },
    {
        id: "unique_list2",
        name: "list2",
        user_id: "user2"
    }
];

const UserLists = () => {
    const renderItem: ListRenderItem<List> = ({ item }) => (
        <ListComponent id={item.id} name={item.name} user_id={item.user_id}/>
    );

    return (
        <FlatList data={listData} renderItem={renderItem} keyExtractor={list => list.id}/>
    );
}

export default UserLists;