/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package ufg.inf.cc.pfc1.blockchain.hyperledgerFabric;

import java.security.PrivateKey;
import java.security.cert.X509Certificate;
import java.util.Set;
import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric.sdk.User;

public class MyUser implements User {
    private String name;
    private String mspId;
    private String affiliation;
    private X509Certificate certificate;

    public MyUser(String name, String affiliation, String mspId, X509Certificate certificate) {
        this.name = name;
        this.affiliation = affiliation;
        this.mspId = mspId;
        this.certificate = certificate;
    }

    @Override
    public String getName() {
        return this.name;
    }

    @Override
    public Set<String> getRoles() {
        return null;  // Implementar conforme necessário
    }

    @Override
    public String getAccount() {
        return null;  // Implementar conforme necessário
    }

    @Override
    public String getAffiliation() {
        return this.affiliation;
    }

    @Override
    public Enrollment getEnrollment() {
        return new Enrollment() {
            @Override
            public PrivateKey getKey() {
                return null;  // Implementar conforme necessário
            }

            @Override
            public String getCert() {
                return certificate.toString();  // Ajustar conforme necessário
            }
        };
    }

    @Override
    public String getMspId() {
        return this.mspId;
    }
}
