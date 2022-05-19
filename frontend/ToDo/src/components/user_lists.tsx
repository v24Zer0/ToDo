import React, { useState } from "react";
import { FlatList, ListRenderItem, Modal, View } from "react-native";
import List from "../models/list";
import ListComponent from "./list_component";
import ListModal from "./list_modal";

const mockLists: List[] = [
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

interface Props {
    navigate(list: List): void;
}

// Add Modal and pass state functions to ListComponent
const UserLists: React.FC<Props> = ({ navigate }) => {
    const [lists, setLists] = useState<List[]>(mockLists);

    const [modalVisible, setModalVisible] = useState<boolean>(false);
    const [modalList, setModalList] = useState<List>({id: "", name: "", user_id: "user1"});

    const renderItem: ListRenderItem<List> = ({ item }) => (
        <ListComponent list={item} setModalVisible={setModalVisible} setModalList={setModalList} navigate={navigate} />
    );

    return (
        <View>
            <Modal visible={modalVisible} onRequestClose={() => setModalVisible(false)}>
                <ListModal list={modalList} />
            </Modal>
            <FlatList data={lists} renderItem={renderItem} keyExtractor={list => list.id}/>
        </View>
    );
}

export default UserLists;