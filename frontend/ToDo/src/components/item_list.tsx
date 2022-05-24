import React, { useEffect, useState } from "react";
import { Button, FlatList, ListRenderItem, Modal, Text, View } from "react-native";
import Item from "../models/item";
import List from "../models/list";
import ItemComponent from "./item_component";
import ItemModal from "./item_modal";

const mockItems: Item[] = [
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

const mockItems2: Item[] = [
    {
        id: "item_id4",
        task: "task4",
        priority: 0,
        list_id: "list_id4"
    },
    {
        id: "item_id5",
        task: "task5",
        priority: 1,
        list_id: "list_id5"
    },
    {
        id: "item_id6",
        task: "random task2",
        priority: 1,
        list_id: "list_id6"
    }
];

interface ItemListProps {
    list: List;
}

// Add Modal and pass setModalVisible, setModalItem to ItemComponent
const ItemList: React.FC<ItemListProps> = ({ list }) => {
    const [items, setItems] = useState<Item[]>(mockItems);
    const [update, setUpdate] = useState<boolean>(false);

    const [modalVisible, setModalVisible] = useState<boolean>(false);
    const [modalItem, setModalItem] = useState<Item>({id: "", list_id: "", task: "", priority: 0});

    const renderItem: ListRenderItem<Item> = ({ item }) => (
        <ItemComponent item={item} setModalVisible={setModalVisible} setModalItem={setModalItem} />
    );

    useEffect(() => {
        update ? setItems(mockItems2) : setItems(mockItems)
    }, [update]);

    return (
        <View>
            <Modal visible={modalVisible} onRequestClose={() => setModalVisible(false)}>
                <ItemModal item={modalItem} list={list} setUpdate={() => setUpdate(!update)} />
            </Modal>
            <FlatList data={items} renderItem={renderItem} keyExtractor={item => item.id}/>
            <Button title="Create new item" onPress={() => { setUpdate(!update) }} />
        </View>
    );
}

export default ItemList;